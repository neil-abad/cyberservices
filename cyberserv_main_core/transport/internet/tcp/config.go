// +build !confonly

package tcp

import (
	"cyberservices.com/core/common"
	"cyberservices.com/core/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
