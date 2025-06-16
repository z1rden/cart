package service_provider

import (
	"cart/internal/cart/clients/product_service"
)

type client struct {
	productService product_service.Client
}

func (s *ServiceProvider) GetProductService() product_service.Client {
	if s.client.productService == nil {
		s.client.productService = product_service.NewClient()
	}

	return s.client.productService
}
