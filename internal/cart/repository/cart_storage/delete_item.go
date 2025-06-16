package cart_storage

import (
	"context"
)

func (s *storage) DeleteItem(ctx context.Context, userID int64, skuID int64) error {
	s.Lock()
	defer s.Unlock()

	cart, exists := cartStorage[userID]
	if !exists {
		return nil
	}

	_, exists = cart.items[skuID]
	if !exists {
		return nil
	}

	delete(cart.items, skuID)

	return nil
}
