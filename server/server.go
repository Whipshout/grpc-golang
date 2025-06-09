package main

import (
	pb "github.com/whipshout/grpc/proto/todo/v1"
)

type server struct {
	d Db
	pb.UnimplementedTodoServiceServer
}
