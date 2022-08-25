// +build !linux,!freebsd
// +build !confonly

package tcp

import (
	"cyberservices.com/core/common/net"
	"cyberservices.com/core/transport/internet"
)

func GetOriginalDestination(conn internet.Connection) (net.Destination, error) {
	return net.Destination{}, nil
}
