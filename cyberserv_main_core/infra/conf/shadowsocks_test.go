package conf_test

import (
	"testing"

	"cyberservices.com/core/common/net"
	"cyberservices.com/core/common/protocol"
	"cyberservices.com/core/common/serial"
	. "cyberservices.com/core/infra/conf"
	"cyberservices.com/core/proxy/shadowsocks"
)

func TestShadowsocksServerConfigParsing(t *testing.T) {
	creator := func() Buildable {
		return new(ShadowsocksServerConfig)
	}

	runMultiTestCase(t, []TestCase{
		{
			Input: `{
				"method": "aes-128-cfb",
				"password": "Project CS-password"
			}`,
			Parser: loadJSON(creator),
			Output: &shadowsocks.ServerConfig{
				User: &protocol.User{
					Account: serial.ToTypedMessage(&shadowsocks.Account{
						CipherType: shadowsocks.CipherType_AES_128_CFB,
						Password:   "Project CS-password",
					}),
				},
				Network: []net.Network{net.Network_TCP},
			},
		},
	})
}
