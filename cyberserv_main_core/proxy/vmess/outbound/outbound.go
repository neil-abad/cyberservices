// +build !confonly

package outbound

//go:generate go run cyberservices.com/core/common/errors/errorgen

import (
	"context"
	"time"

	"cyberservices.com/core"
	"cyberservices.com/core/common"
	"cyberservices.com/core/common/buf"
	"cyberservices.com/core/common/net"
	"cyberservices.com/core/common/platform"
	"cyberservices.com/core/common/protocol"
	"cyberservices.com/core/common/retry"
	"cyberservices.com/core/common/session"
	"cyberservices.com/core/common/signal"
	"cyberservices.com/core/common/task"
	"cyberservices.com/core/features/policy"
	"cyberservices.com/core/proxy/vmess"
	"cyberservices.com/core/proxy/vmess/encoding"
	"cyberservices.com/core/transport"
	"cyberservices.com/core/transport/internet"
)

// Handler is an outbound connection handler for VMess protocol.
type Handler struct {
	serverList    *protocol.ServerList
	serverPicker  protocol.ServerPicker
	policyManager policy.Manager
}

// New creates a new VMess outbound handler.
func New(ctx context.Context, config *Config) (*Handler, error) {
	serverList := protocol.NewServerList()
	for _, rec := range config.Receiver {
		s, err := protocol.NewServerSpecFromPB(rec)
		if err != nil {
			return nil, newError("failed to parse server spec").Base(err)
		}
		serverList.AddServer(s)
	}

	v := core.MustFromContext(ctx)
	handler := &Handler{
		serverList:    serverList,
		serverPicker:  protocol.NewRoundRobinServerPicker(serverList),
		policyManager: v.GetFeature(policy.ManagerType()).(policy.Manager),
	}

	return handler, nil
}

// Process implements proxy.Outbound.Process().
func (h *Handler) Process(ctx context.Context, link *transport.Link, dialer internet.Dialer) error {
	var rec *protocol.ServerSpec
	var conn internet.Connection

	err := retry.ExponentialBackoff(5, 200).On(func() error {
		rec = h.serverPicker.PickServer()
		rawConn, err := dialer.Dial(ctx, rec.Destination())
		if err != nil {
			return err
		}
		conn = rawConn

		return nil
	})
	if err != nil {
		return newError("failed to find an available destination").Base(err).AtWarning()
	}
	defer conn.Close() //nolint: errcheck

	outbound := session.OutboundFromContext(ctx)
	if outbound == nil || !outbound.Target.IsValid() {
		return newError("target not specified").AtError()
	}

	target := outbound.Target
	newError("tunneling request to ", target, " via ", rec.Destination()).WriteToLog(session.ExportIDToError(ctx))

	command := protocol.RequestCommandTCP
	if target.Network == net.Network_UDP {
		command = protocol.RequestCommandUDP
	}
	if target.Address.Family().IsDomain() && target.Address.Domain() == "v1.mux.cool" {
		command = protocol.RequestCommandMux
	}

	user := rec.PickUser()
	request := &protocol.RequestHeader{
		Version: encoding.Version,
		User:    user,
		Command: command,
		Address: target.Address,
		Port:    target.Port,
		Option:  protocol.RequestOptionChunkStream,
	}

	account := request.User.Account.(*vmess.MemoryAccount)
	request.Security = account.Security

	if request.Security == protocol.SecurityType_AES128_GCM || request.Security == protocol.SecurityType_NONE || request.Security == protocol.SecurityType_CHACHA20_POLY1305 {
		request.Option.Set(protocol.RequestOptionChunkMasking)
	}

	if shouldEnablePadding(request.Security) && request.Option.Has(protocol.RequestOptionChunkMasking) {
		request.Option.Set(protocol.RequestOptionGlobalPadding)
	}

	input := link.Reader
	output := link.Writer

	isAEAD := false
	if !aead_disabled && len(account.AlterIDs) == 0 {
		isAEAD = true
	}

	session := encoding.NewClientSession(isAEAD, protocol.DefaultIDHash, ctx)
	sessionPolicy := h.policyManager.ForLevel(request.User.Level)

	ctx, cancel := context.WithCancel(ctx)
	timer := signal.CancelAfterInactivity(ctx, cancel, sessionPolicy.Timeouts.ConnectionIdle)

	requestDone := func() error {
		defer timer.SetTimeout(sessionPolicy.Timeouts.DownlinkOnly)

		writer := buf.NewBufferedWriter(buf.NewWriter(conn))
		if err := session.EncodeRequestHeader(request, writer); err != nil {
			return newError("failed to encode request").Base(err).AtWarning()
		}

		bodyWriter := session.EncodeRequestBody(request, writer)
		if err := buf.CopyOnceTimeout(input, bodyWriter, time.Millisecond*100); err != nil && err != buf.ErrNotTimeoutReader && err != buf.ErrReadTimeout {
			return newError("failed to write first payload").Base(err)
		}

		if err := writer.SetBuffered(false); err != nil {
			return err
		}

		if err := buf.Copy(input, bodyWriter, buf.UpdateActivity(timer)); err != nil {
			return err
		}

		if request.Option.Has(protocol.RequestOptionChunkStream) {
			if err := bodyWriter.WriteMultiBuffer(buf.MultiBuffer{}); err != nil {
				return err
			}
		}

		return nil
	}

	responseDone := func() error {
		defer timer.SetTimeout(sessionPolicy.Timeouts.UplinkOnly)

		reader := &buf.BufferedReader{Reader: buf.NewReader(conn)}
		header, err := session.DecodeResponseHeader(reader)
		if err != nil {
			return newError("failed to read header").Base(err)
		}
		h.handleCommand(rec.Destination(), header.Command)

		bodyReader := session.DecodeResponseBody(request, reader)

		return buf.Copy(bodyReader, output, buf.UpdateActivity(timer))
	}

	var responseDonePost = task.OnSuccess(responseDone, task.Close(output))
	if err := task.Run(ctx, requestDone, responseDonePost); err != nil {
		return newError("connection ends").Base(err)
	}

	return nil
}

var (
	enablePadding = false
	aead_disabled = false
)

func shouldEnablePadding(s protocol.SecurityType) bool {
	return enablePadding || s == protocol.SecurityType_AES128_GCM || s == protocol.SecurityType_CHACHA20_POLY1305 || s == protocol.SecurityType_AUTO
}

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return New(ctx, config.(*Config))
	}))

	const defaultFlagValue = "NOT_DEFINED_AT_ALL"

	paddingValue := platform.NewEnvFlag("cyberservices.vmess.padding").GetValue(func() string { return defaultFlagValue })
	if paddingValue != defaultFlagValue {
		enablePadding = true
	}

	aeadDisabled := platform.NewEnvFlag("cyberservices.vmess.aead.disabled").GetValue(func() string { return defaultFlagValue })
	if aeadDisabled == "true" {
		aead_disabled = true
	}
}
