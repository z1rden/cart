package loms_service

import (
	"cart/pkg/api/order"
	"context"
	"google.golang.org/grpc"
)

type Client interface {
	OrderCreate(ctx context.Context, user int64, items []*OrderItem) (int64, error)
}

type client struct {
	orderGrpcClient order.OrderClient
}

func NewClient(conn *grpc.ClientConn) Client {
	return &client{
		orderGrpcClient: order.NewOrderClient(conn),
	}
}
