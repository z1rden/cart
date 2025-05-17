package cart_service_test

import (
	"cart/internal/cart/client/product_service"
	"cart/internal/cart/service/cart_service"
	"cart/internal/cart/suite"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestServiceAddItem(t *testing.T) {
	type test struct {
		Name                string
		UserID              int64
		SkuID               int64
		Count               uint16
		Product             *product_service.Product
		ClientError         error
		StorageAddItemError error
		Error               error
	}

	var (
		StorageAddItemError = errors.New("Возникла проблема с добавлением в Storage")
		ClientError         = errors.New("Возникла проблема при получении товара из другого сервиса")
	)

	var tests = []test{
		{
			Name:        "Отсутствие товара в productService",
			UserID:      1,
			SkuID:       1,
			Count:       1,
			ClientError: ClientError,
			Error:       ClientError,
		},
		{
			Name:                "Ошибка при добавлении в Storage",
			UserID:              1,
			SkuID:               1,
			Count:               1,
			StorageAddItemError: StorageAddItemError,
			Error:               StorageAddItemError,
		},
		{
			Name:   "Успешный тест",
			UserID: 1,
			SkuID:  1,
			Count:  1,
			Product: &product_service.Product{
				Name:  "example1",
				Price: 10,
			},
		},
	}

	t.Parallel()

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			sp := suite.NewSuiteProvider()
			s := cart_service.NewService(
				sp.GetCartStorageMock(),
				sp.GetProductServiceMock(),
				sp.GetLomsServiceMock(),
			)
			ctx := context.Background()

			sp.GetProductServiceMock().EXPECT().
				GetProduct(mock.Anything, tt.SkuID).
				Return(tt.Product, tt.ClientError)

			sp.GetCartStorageMock().EXPECT().
				AddItem(mock.Anything, tt.UserID, tt.SkuID, tt.Count).
				Return(tt.StorageAddItemError)

			err := s.AddItem(ctx, tt.UserID, tt.SkuID, tt.Count)
			if tt.Error != nil {
				assert.NotNil(t, tt.Error, "Должна возникать ошибка")
				assert.ErrorIs(t, err, tt.Error, "Должна соответствовать этой ошибке")
			} else {
				assert.NoError(t, err, "Ошибки не должно возникать")
			}
		})
	}
}
