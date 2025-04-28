package cart_storage

import (
	"context"
	"sync"
)

type Storage interface {
	AddItem(ctx context.Context, userID int64, skuID int64, count uint16) error
	DeleteItem(ctx context.Context, userID int64, skuID int64) error
	GetItemsByUserID(ctx context.Context, userID int64) (*Cart, error)
	DeleteItemsByUserID(ctx context.Context, userID int64) error
}

var cartStorage = map[int64]*Cart{}

type storage struct {
	sync.RWMutex
}

func NewStorage() Storage {
	return &storage{}
}
