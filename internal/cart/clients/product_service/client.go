package product_service

import (
	"context"
	"errors"
)

type Client interface {
	GetProduct(ctx context.Context, skuID int64) error
}

type client struct {
	storage map[int64]uint16
}

func NewClient() Client {
	return &client{
		storage: map[int64]uint16{
			1: 10,
			2: 20,
			3: 30,
			4: 40,
			5: 50,
		},
	}
}

func (c *client) GetProduct(ctx context.Context, skuID int64) error {
	_, exists := c.storage[skuID]
	if !exists {
		return errors.New("sku not found")
	}

	return nil
}
