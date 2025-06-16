package cart_storage

import (
	"context"
)

func (s *storage) DeleteItemsByUserID(ctx context.Context, userID int64) error {
	s.Lock()
	defer s.Unlock()

	_, exists := cartStorage[userID]
	if exists {
		delete(cartStorage, userID)
	}

	return nil
}
