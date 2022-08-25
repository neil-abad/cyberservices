package udp

import (
	"cyberservices.com/core/common"
	"cyberservices.com/core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
