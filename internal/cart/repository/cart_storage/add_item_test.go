package cart_storage

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorageAddItem(t *testing.T) {
	ctx := context.Background()
	s := NewStorage()

	err := s.AddItem(ctx, 1, 1, 1)
	assert.NoError(t, err, "При добавлении в корзину продукта ошибки не должно возникать")
}
