package productusecase

import (
	"github.com/mamazinho/go-clean-architeture/core/domain"
	"github.com/mamazinho/go-clean-architeture/core/dto"
)

func (usecase usecase) Fetch(paginationRequest *dto.PaginationRequestParms) (*domain.Pagination[[]domain.Product], error) {
	products, err := usecase.repository.Fetch(paginationRequest)

	if err != nil {
		return nil, err
	}

	return products, nil
}
	