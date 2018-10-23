package main

import (
	"flag"
	"fmt"
	pb "github.com/nokamoto/proposal-nginx-grpc/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

var (
	port = flag.Int("p", 9000, "grpc server port")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		panic(fmt.Sprintf("listen tcp port (%d) - %v", *port, err))
	}

	fmt.Printf("listen tcp port (%d)\n", *port)

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)

	svc := &service{}

	pb.RegisterPingServiceServer(server, svc)
	reflection.Register(server)

	err = server.Serve(lis)
	if err != nil {
		panic(fmt.Sprintf("serve %v - %v", lis, err))
	}
}
