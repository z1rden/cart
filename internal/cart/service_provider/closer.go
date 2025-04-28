package service_provider

import (
	"cart/internal/cart/closer"
	"os"
	"syscall"
)

func (s *ServiceProvider) GetCloser() closer.Closer {
	if s.closer == nil {
		s.closer = closer.NewCloser(os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	}

	return s.closer
}
