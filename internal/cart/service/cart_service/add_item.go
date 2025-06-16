package cart_service

import (
	"cart/internal/cart/logger"
	"context"
)

func (s *service) AddItem(ctx context.Context, userID int64, skuID int64, count uint16) error {
	const operation = "cart_service.AddItem"

	_, err := s.productService.GetProduct(ctx, skuID)
	if err != nil {
		logger.Errorf(ctx, "%s: failed to get product: %v", operation, err)

		return err
	}

	err = s.cartStorage.AddItem(ctx, userID, skuID, count)
	if err != nil {
		logger.Errorf(ctx, "%s: failed to add item: %v", operation, err)

		return err
	}

	return nil
}
