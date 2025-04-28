package cart_api

import "net/http"

type HttpAPIHandler struct {
	Pattern string
	Handler http.HandlerFunc
}
