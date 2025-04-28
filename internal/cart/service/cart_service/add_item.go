package cart_service

import "context"

func (s *service) AddItem(ctx context.Context, userID int64, skuID int64, count uint16) error {
	_, err := s.productService.GetProduct(ctx, skuID)

	if err != nil {
		return err
	}

	err = s.cartStorage.AddItem(ctx, userID, skuID, count)
	if err != nil {
		return err
	}

	return nil
}
