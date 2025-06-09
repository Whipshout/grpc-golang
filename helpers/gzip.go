package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	pb "github.com/whipshout/grpc/helpers/proto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

func main() {
	task := &pb.Task{
		Id:          1,
		Description: "This is a task",
		DueDate:     timestamppb.New(time.Now().Add(5 * 24 * time.Hour)),
	}

	o, c := compressedSize(task)

	fmt.Printf("original: %d\ncompressed: %d\n", o, c)
}

func compressedSize[M protoreflect.ProtoMessage](msg M) (int, int) {
	var b bytes.Buffer

	gz := gzip.NewWriter(&b)

	out, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := gz.Write(out); err != nil {
		log.Fatal(err)
	}

	if err := gz.Close(); err != nil {
		log.Fatal(err)
	}

	return len(out), len(b.Bytes())
}
