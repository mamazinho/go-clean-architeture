package productrepository

import (
	"github.com/mamazinho/go-clean-architeture/adapter/postgres"
	"github.com/mamazinho/go-clean-architeture/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

// New returns contract implementation of ProductRepository
func New(db postgres.PoolInterface) domain.ProductRepository {
	return &repository{
		db: db,
	}
}
