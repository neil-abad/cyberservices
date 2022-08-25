package command_test

import (
	"context"
	"testing"

	"cyberservices.com/core"
	"cyberservices.com/core/app/dispatcher"
	"cyberservices.com/core/app/log"
	. "cyberservices.com/core/app/log/command"
	"cyberservices.com/core/app/proxyman"
	_ "cyberservices.com/core/app/proxyman/inbound"
	_ "cyberservices.com/core/app/proxyman/outbound"
	"cyberservices.com/core/common"
	"cyberservices.com/core/common/serial"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
