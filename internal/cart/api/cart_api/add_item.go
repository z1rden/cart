package cart_api

import (
	"cart/internal/cart/logger"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func (a *api) AddItem() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		req, err := toAddItemRequest(ctx, r)
		if err != nil {
			logger.Errorf(ctx, "handleAddItem: request is not valid: %v", err)
			http.Error(w, fmt.Sprintf("request is not valid: %s", err), http.StatusBadRequest)
			return
		}

		err = a.cartService.AddItem(r.Context(), req.UserID, req.SkuID, req.Quantity)
		if err != nil {
			logger.Errorf(ctx, "handleAddItem failed to add item: %v", err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func toAddItemRequest(ctx context.Context, r *http.Request) (*AddItemRequest, error) {
	userID, _ := strconv.ParseInt(r.PathValue("user_id"), 10, 64)
	skuID, _ := strconv.ParseInt(r.PathValue("sku_id"), 10, 64)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	data := &AddItemRequestBody{}
	err = json.Unmarshal(body, data)
	if err != nil {
		logger.Errorf(ctx, "handleAddItem: failed to unmarshal body json: %v", err)
		return nil, err
	}

	req := &AddItemRequest{
		UserID:   userID,
		SkuID:    skuID,
		Quantity: data.Count,
	}

	return req, nil
}
