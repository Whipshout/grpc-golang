package main

import (
	pb "github.com/whipshout/grpc/proto/todo/v1"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatalln("usage: server [IP_ADDR]")
	}

	addr := args[0]

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	defer func(lis net.Listener) {
		if err := lis.Close(); err != nil {
			log.Fatalf("unexpected error: %v\n", err)
		}
	}(listener)

	log.Printf("listening at %s\n", addr)

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	defer s.Stop()

	pb.RegisterTodoServiceServer(s, &server{
		d: New(),
	})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
