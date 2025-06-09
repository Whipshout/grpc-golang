package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"log"

	pb "github.com/whipshout/grpc/helpers/proto"
)

func main() {
	t := &pb.Tags{}

	tags := []int{1, 16, 2048, 262_144, 33_554_432, 33_554_432, 536_870_911}

	fields := []*int32{&t.Tag, &t.Tag2, &t.Tag3, &t.Tag4, &t.Tag5, &t.Tag6}

	sz := serializedSizePb(t)
	fmt.Printf("0 - %d\n", sz)

	for i, f := range fields {
		*f = 1

		sz := serializedSizePb(t)
		fmt.Printf("%d - %d\n", tags[i], sz-(i+1))
	}
}

func serializedSizePb[M protoreflect.ProtoMessage](msg M) int {
	out, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	return len(out)
}
