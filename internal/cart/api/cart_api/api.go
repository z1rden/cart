package cart_api

import (
	"cart/internal/cart/service/cart_service"
	"fmt"
	"net/http"
)

type API interface {
	GetHandlers() []HttpAPIHandler
	AddItem() func(http.ResponseWriter, *http.Request)
}

type api struct {
	cartService cart_service.Service
}

func NewApi(cartService cart_service.Service) API {
	return &api{
		cartService: cartService,
	}
}

func (a *api) GetHandlers() []HttpAPIHandler {
	return []HttpAPIHandler{
		{
			Pattern: fmt.Sprintf("%s /user/{%s}/cart/{%s}", http.MethodPost, "user_id", "sku_id"),
			Handler: a.AddItem(),
		},
	}
}
