package cart_api

import "testing"

func TestAddItem(t *testing.T) {
	type test struct {
		Name     string
		UserID   int64
		SkuID    int64
		Quantity uint16
	}

	tests := []*test{
		{
			Name:     "Не передан пользователь",
			UserID:   0,
			SkuID:    1,
			Quantity: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {

		})
	}
}
