package suite

import (
	"cart/internal/cart/clients/product_service"
	product_service_mock "cart/internal/cart/clients/product_service/mocks"
	"cart/internal/cart/repository/cart_storage"
	cart_storage_mock "cart/internal/cart/repository/cart_storage/mocks"
)

type suiteProvider struct {
	cartStorage        cart_storage.Storage
	cartStorageMock    *cart_storage_mock.StorageMock
	productService     product_service.Client
	productServiceMock *product_service_mock.ClientMock
}

func NewSuiteProvider() *suiteProvider {
	return &suiteProvider{}
}

func (s *suiteProvider) GetCartStorageMock() *cart_storage_mock.StorageMock {
	if s.cartStorageMock == nil {
		s.cartStorageMock = &cart_storage_mock.StorageMock{}
	}

	return s.cartStorageMock
}

func (s *suiteProvider) GetCartStorage() cart_storage.Storage {
	if s.cartStorage == nil {
		s.cartStorage = s.GetCartStorageMock()
	}

	return s.cartStorage
}

func (s *suiteProvider) GetProductServiceMock() *product_service_mock.ClientMock {
	if s.productServiceMock == nil {
		s.productServiceMock = &product_service_mock.ClientMock{}
	}

	return s.productServiceMock
}

func (s *suiteProvider) GetProductService() product_service.Client {
	if s.productService == nil {
		s.productService = s.GetProductServiceMock()
	}

	return s.productService
}
