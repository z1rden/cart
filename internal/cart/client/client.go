package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetClientConn(ctx context.Context, port string) *grpc.ClientConn {
	conn, err := grpc.NewClient(fmt.Sprintf(":%s", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return conn
}
