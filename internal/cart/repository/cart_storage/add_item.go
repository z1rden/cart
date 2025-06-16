package cart_storage

import "context"

func (s *storage) AddItem(ctx context.Context, userID int64, skuID int64, count uint16) error {
	s.Lock()
	defer s.Unlock()

	cart, ok := cartStorage[userID]
	if !ok {
		cart = &Cart{
			items: map[int64]uint16{},
		}
		cartStorage[userID] = cart
	}

	cart.items[skuID] += count

	return nil
}
