package main

import (
	pb "github.com/whipshout/grpc/proto/todo/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	_ "google.golang.org/grpc/encoding/gzip"
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

	creds, err := credentials.NewServerTLSFromFile("./certs/server_cert.pem", "./certs/server_key.pem")
	if err != nil {
		log.Fatalf("failed to generate credentials %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.Creds(creds),
		grpc.ChainUnaryInterceptor(unaryAuthInterceptor, unaryLogInterceptor),
		grpc.ChainStreamInterceptor(streamAuthInterceptor, streamLogInterceptor),
	}
	s := grpc.NewServer(opts...)
	defer s.Stop()

	pb.RegisterTodoServiceServer(s, &server{
		d: New(),
	})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
