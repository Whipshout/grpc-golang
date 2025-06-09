package main

import (
	pb "github.com/whipshout/grpc/proto/todo/v2"
)

type server struct {
	d Db
	pb.UnimplementedTodoServiceServer
}
