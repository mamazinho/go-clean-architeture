package productrepository

import (
	"context"

	"github.com/booscaaa/go-paginate/paginate"
	"github.com/mamazinho/go-clean-architeture/core/domain"
	"github.com/mamazinho/go-clean-architeture/core/dto"
)

// Fetch implements domain.ProductRepository
func (r *repository) Fetch(paginationRequest *dto.PaginationRequestParms) (*domain.Pagination[[]domain.Product], error) {
	ctx := context.Background()
	products := []domain.Product{}
	total := int32(0)

	pagin := paginate.Instance(nil)

	query, queryCount := pagin.
		Query("SELECT * FROM product").
		Page(paginationRequest.Page).
		Desc(paginationRequest.Descending).
		Sort(paginationRequest.Sort).
		RowsPerPage(paginationRequest.ItemsPerPage).
		SearchBy(paginationRequest.Search, "name", "description").
		Select()

	{
		rows, err := r.db.Query(
			ctx,
			*query,
		)

		if err != nil {
			return nil, err
		}

		for rows.Next() {
			product := domain.Product{}

			rows.Scan(
				&product.ID,
				&product.Name,
				&product.Price,
				&product.Description,
			)

			products = append(products, product)
		}
	}

	{
		err := r.db.QueryRow(ctx, *queryCount).Scan(&total)

		if err != nil {
			return nil, err
		}
	}

	return &domain.Pagination[[]domain.Product]{
		Items: products,
		Total: total,
	}, nil

}
