package cart_api

type AddItemRequestBody struct {
	Count uint16 `json:"count"`
}

type AddItemRequest struct {
	UserID   int64  `validate:"required|int|min:1"`
	SkuID    int64  `validate:"required|int|min:1"`
	Quantity uint16 `validate:"required|int|min:1"`
}

type DeleteItemRequest struct {
	UserID int64 `validate:"required|int|min:1"`
	SkuID  int64 `validate:"required|int|min:1"`
}

type GetCartByUserIDRequest struct {
	UserID int64 `validate:"required|int|min:1"`
}

type DeleteItemsByUserIDRequest struct {
	UserID int64 `validate:"required|int|min:1"`
}
