package cart_service

import (
	"cart/internal/cart/clients/product_service"
	"cart/internal/cart/repository/cart_storage"
	"context"
)

type Service interface {
	AddItem(ctx context.Context, userID int64, skuID int64, count uint16) error
	DeleteItem(ctx context.Context, userID int64, skuID int64) error
	DeleteCartByUserId(ctx context.Context, userID int64) error
	GetCartByUserID(ctx context.Context, userID int64) (*Cart, error)
}

type service struct {
	productService product_service.Client
	cartStorage    cart_storage.Storage
}

func NewService(cartStorage cart_storage.Storage, productService product_service.Client) Service {
	return &service{
		cartStorage:    cartStorage,
		productService: productService,
	}
}
