package main

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "github.com/nokamoto/proposal-nginx-grpc/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct{}

func (s *service) Ping(context.Context, *pb.Empty) (*timestamp.Timestamp, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented yet")
}

func (s *service) PingC(pb.PingService_PingCServer) error {
	return status.Error(codes.Unimplemented, "not implemented yet")
}

func (s *service) PingS(*pb.List, pb.PingService_PingSServer) error {
	return status.Error(codes.Unimplemented, "not implemented yet")
}

func (s *service) PingB(pb.PingService_PingBServer) error {
	return status.Error(codes.Unimplemented, "not implemented yet")
}
