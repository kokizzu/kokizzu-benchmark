package main

import (
	"kokizzu-benchmark/grpc-vs-socketio/flatbuffer/schema/users"
	"context"
	"log"
	"net"

	flatbuffers "github.com/google/flatbuffers/go"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Add(context context.Context, request *users.AddRequest) (*flatbuffers.Builder, error) {
	b := flatbuffers.NewBuilder(0)
	users.AddResponseStart(b)
	b.Finish(users.AddRequestEnd(b))

	return b, nil
}

func main() {
	lis, err := net.Listen("tcp", ":15001")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer(grpc.CustomCodec(flatbuffers.FlatbuffersCodec{}))
	users.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("serve: %v", err)
	}
}
