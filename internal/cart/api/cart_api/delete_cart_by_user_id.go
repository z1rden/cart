package cart_api

import (
	"cart/internal/cart/logger"
	"context"
	"fmt"
	"net/http"
	"strconv"
)

func (a *api) DeleteCartByUserID() func(w http.ResponseWriter, r *http.Request) {
	const operation = "api.DeleteCartByUserID"

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userID, err := toDeleteCartByUserIdRequest(ctx, r)
		if err != nil {
			logger.Errorf(ctx, "%s: request is not valid: %v", operation, err)
			http.Error(w, fmt.Sprintf("request is not valid: %v", err), http.StatusBadRequest)

			return
		}

		a.cartService.DeleteCartByUserId(ctx, userID)

		w.WriteHeader(http.StatusNoContent)
	}
}

func toDeleteCartByUserIdRequest(ctx context.Context, r *http.Request) (int64, error) {
	const operatuion = "api.toDeleteCartByUserIDRequest"

	userID, err := strconv.ParseInt(r.PathValue("user_id"), 10, 64)
	if err != nil {
		logger.Errorf(ctx, "%s: userID is not valid: %v", operatuion, err)

		return 0, fmt.Errorf("userID is not valid: %s", err)
	}

	return userID, nil
}
