package server

import (
	"context"
	"errors"
	"fmt"

	"github.com/thanhpk/randstr"

	"github.com/aifaniyi/grpc-sample/pkg/dto"
)

type Server struct {
	dto.GreetingServiceServer
}

func (s *Server) Greeting(ctx context.Context,
	req *dto.GreetingServiceRequest) (*dto.GreetingServiceReply, error) {

	if req == nil {
		return nil, errors.New("request cannot be null")
	}

	return &dto.GreetingServiceReply{
		Message: fmt.Sprintf("Hello, received %d", req.GetCount()),
	}, nil
}

func (s *Server) LotsOfGreetingReplies(req *dto.GreetingServiceRequest, srv dto.GreetingService_LotsOfGreetingRepliesServer) error {

	var i int64
	count := req.GetCount()

	for ; i < count; i++ {

		if err := srv.Send(&dto.GreetingServiceReply{
			Id:      i,
			Message: randstr.String(16),
		}); err != nil {
			return err
		}
	}

	return nil
}
