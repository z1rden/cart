package cart_storage

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorageDeleteCartByUserID(t *testing.T) {
	ctx := context.Background()
	s := NewStorage()

	err := s.DeleteCartByUserID(ctx, 1)
	assert.NoError(t, err, "Ошибки не должно вернуться по условию задачи")

	_ = s.AddItem(ctx, 1, 1, 1)
	_ = s.AddItem(ctx, 1, 2, 5)
	err = s.DeleteCartByUserID(ctx, 1)
	assert.NoError(t, err, "Ошибки не должно вернуться при удалении непустой корзины")
}
