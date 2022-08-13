package routing

import (
	"context"

	"v2ray.com/core/common/net"
	"v2ray.com/core/features"
	"v2ray.com/core/transport"
)

//
// :api:stable
type Dispatcher interface {
	features.Feature

	// Dispatch returns a Ray for transporting data for the given request.
	Dispatch(ctx context.Context, dest net.Destination) (*transport.Link, error)
}

// DispatcherType returns the type of Dispatcher interface. Can be used to implement common.HasType.
//
// :api:stable
func DispatcherType() interface{} {
	return (*Dispatcher)(nil)
}
