package mux_test

import (
	"testing"

	"cyberservices.com/core/common"
	"cyberservices.com/core/common/buf"
	"cyberservices.com/core/common/mux"
	"cyberservices.com/core/common/net"
)

func BenchmarkFrameWrite(b *testing.B) {
	frame := mux.FrameMetadata{
		Target:        net.TCPDestination(net.DomainAddress("www.cyberservices.com"), net.Port(80)),
		SessionID:     1,
		SessionStatus: mux.SessionStatusNew,
	}
	writer := buf.New()
	defer writer.Release()

	for i := 0; i < b.N; i++ {
		common.Must(frame.WriteTo(writer))
		writer.Clear()
	}
}
