package productservice

import (
	"github.com/mamazinho/go-clean-architeture/core/domain"
)

type service struct {
	usecase domain.ProductUseCase
}

// New returns contract implementation of ProductService
func New(usecase domain.ProductUseCase) domain.ProductService {
	return &service{
		usecase: usecase,
	}
}
