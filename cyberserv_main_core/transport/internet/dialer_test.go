package internet_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"

	"cyberservices.com/core/common"
	"cyberservices.com/core/common/net"
	"cyberservices.com/core/testing/servers/tcp"
	. "cyberservices.com/core/transport/internet"
)

func TestDialWithLocalAddr(t *testing.T) {
	server := &tcp.Server{}
	dest, err := server.Start()
	common.Must(err)
	defer server.Close()

	conn, err := DialSystem(context.Background(), net.TCPDestination(net.LocalHostIP, dest.Port), nil)
	common.Must(err)
	if r := cmp.Diff(conn.RemoteAddr().String(), "127.0.0.1:"+dest.Port.String()); r != "" {
		t.Error(r)
	}
	conn.Close()
}
