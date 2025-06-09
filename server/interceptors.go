package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
)

const authTokenKey = "auth_token"
const authTokenValue = "authd"

func validateAuthToken(ctx context.Context) error {
	md, _ := metadata.FromIncomingContext(ctx)
	if t, ok := md[authTokenKey]; ok {
		switch {
		case len(t) != 1:
			return status.Errorf(codes.InvalidArgument, "%s should contain only 1 value", authTokenKey)
		case t[0] != authTokenValue:
			return status.Errorf(codes.Unauthenticated, "incorrect %s", authTokenKey)
		}
	} else {
		return status.Errorf(codes.Unauthenticated, "missing %s", authTokenKey)
	}

	return nil
}

func unaryAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if err := validateAuthToken(ctx); err != nil {
		return nil, err
	}

	return handler(ctx, req)
}

func streamAuthInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if err := validateAuthToken(ss.Context()); err != nil {
		return err
	}

	return handler(srv, ss)
}

func unaryLogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println(info.FullMethod, "called")
	return handler(ctx, req)
}

func streamLogInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println(info.FullMethod, "called")
	return handler(srv, ss)
}
