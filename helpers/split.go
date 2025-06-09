package main

import (
	"fmt"
	pb "github.com/whipshout/grpc/helpers/proto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"log"
)

func main() {
	s := &pb.Split{
		Name: "Packt",
	}

	sz := serializedSizeSplit(s)

	fmt.Printf("With name: %d\n", sz)

	s.Name = ""
	s.ComplexName = &pb.ComplexName{Name: "Packt"}
	sz = serializedSizeSplit(s)

	fmt.Printf("With ComplexName: %d\n", sz)
}

func serializedSizeSplit[M protoreflect.ProtoMessage](msg M) int {
	out, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	return len(out)
}
