package core_test

import (
	"testing"

	proto "github.com/golang/protobuf/proto"
	. "cyberservices.com/core"
	"cyberservices.com/core/app/dispatcher"
	"cyberservices.com/core/app/proxyman"
	"cyberservices.com/core/common"
	"cyberservices.com/core/common/net"
	"cyberservices.com/core/common/protocol"
	"cyberservices.com/core/common/serial"
	"cyberservices.com/core/common/uuid"
	"cyberservices.com/core/features/dns"
	"cyberservices.com/core/features/dns/localdns"
	_ "cyberservices.com/core/main/distro/all"
	"cyberservices.com/core/proxy/dokodemo"
	"cyberservices.com/core/proxy/vmess"
	"cyberservices.com/core/proxy/vmess/outbound"
	"cyberservices.com/core/testing/servers/tcp"
)

func TestProject CSDependency(t *testing.T) {
	instance := new(Instance)

	wait := make(chan bool, 1)
	instance.RequireFeatures(func(d dns.Client) {
		if d == nil {
			t.Error("expected dns client fulfilled, but actually nil")
		}
		wait <- true
	})
	instance.AddFeature(localdns.New())
	<-wait
}

func TestProject CSClose(t *testing.T) {
	port := tcp.PickPort()

	userId := uuid.New()
	config := &Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
		Inbound: []*InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortRange: net.SinglePortRange(port),
					Listen:    net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address: net.NewIPOrDomain(net.LocalHostIP),
					Port:    uint32(0),
					NetworkList: &net.NetworkList{
						Network: []net.Network{net.Network_TCP, net.Network_UDP},
					},
				}),
			},
		},
		Outbound: []*OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&outbound.Config{
					Receiver: []*protocol.ServerEndpoint{
						{
							Address: net.NewIPOrDomain(net.LocalHostIP),
							Port:    uint32(0),
							User: []*protocol.User{
								{
									Account: serial.ToTypedMessage(&vmess.Account{
										Id: userId.String(),
									}),
								},
							},
						},
					},
				}),
			},
		},
	}

	cfgBytes, err := proto.Marshal(config)
	common.Must(err)

	server, err := StartInstance("protobuf", cfgBytes)
	common.Must(err)
	server.Close()
}
