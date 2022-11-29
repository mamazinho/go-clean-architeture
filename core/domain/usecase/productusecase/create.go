package productusecase

import (
	"github.com/mamazinho/go-clean-architeture/core/domain"
	"github.com/mamazinho/go-clean-architeture/core/dto"
)

func (usecase usecase) Create(productRequest *dto.CreateProductRequest) (*domain.Product, error) {
	product, err := usecase.repository.Create(productRequest)

	if err != nil {
		return nil, err
	}

	return product, nil
}
