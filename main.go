package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "example.com/m/protobuf"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 5000, "The server port")
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayMsg(ctx context.Context, in *pb.MsgRequest) (*pb.MsgReply, error) {
	return &pb.MsgReply{Message: "Hello again " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
