package productusecase

import (
	"github.com/mamazinho/go-clean-architeture/core/domain"
)

type usecase struct {
	repository domain.ProductRepository
}

// New returns contract implementation of ProductUseCase
func New(repository domain.ProductRepository) domain.ProductUseCase {
	return &usecase{
		repository: repository,
	}
}
