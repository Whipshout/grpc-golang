package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"log"
	"unsafe"
)

func main() {
	var data uint64 = 72_057_594_037_927_936

	ui64 := &wrapperspb.UInt64Value{
		Value: data,
	}

	d, w := serializedSize(data, ui64)

	fmt.Printf("in memory: %d\npb: %d\n", d, w)
}

func serializedSize[D constraints.Integer, W protoreflect.ProtoMessage](data D, wrapper W) (uintptr, int) {
	out, err := proto.Marshal(wrapper)
	if err != nil {
		log.Fatal(err)
	}

	return unsafe.Sizeof(data), len(out) - 1
}
