package cart_storage

import (
	"context"
	"errors"
)

func (s *storage) GetItemsByUserID(ctx context.Context, userID int64) (*Cart, error) {
	s.RLock()
	defer s.RUnlock()

	cart, exists := cartStorage[userID]
	if !exists {
		return nil, errors.New("cart for user not found")
	}

	return cart, nil
}
