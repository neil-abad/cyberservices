package udp

import (
	"cyberservices.com/core/common/buf"
	"cyberservices.com/core/common/net"
)

// Packet is a UDP packet together with its source and destination address.
type Packet struct {
	Payload *buf.Buffer
	Source  net.Destination
	Target  net.Destination
}
