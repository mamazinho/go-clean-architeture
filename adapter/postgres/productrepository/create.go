package productrepository

import (
	"context"

	"github.com/mamazinho/go-clean-architeture/core/domain"
	"github.com/mamazinho/go-clean-architeture/core/dto"
)

// Create implements domain.ProductRepository
func (r *repository) Create(productRequest *dto.CreateProductRequest) (*domain.Product, error) {
	ctx := context.Background()
	product := domain.Product{}

	err := r.db.QueryRow(
		ctx,
		"INSERT INTO product (name, price, description) VALUES ($1, $2, $3) returning *",
		productRequest.Name,
		productRequest.Price,
		productRequest.Description,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Description,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
