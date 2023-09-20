// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package pingpong

import (
	"context"
)

// implement the service handler for PingPong service
type PingPongServiceImpl struct{}

func (s *PingPongServiceImpl) Ping(ctx context.Context, in *PingRequest) (*PongReply, error) {
	return &PongReply{Message: "Hello " + in.GetSender()}, nil
}

func (s *PingPongServiceImpl) NoPing(ctx context.Context, in *Empty) (*PongReply, error) {
	return &PongReply{Message: "HelloPong"}, nil
}
