package mysql

import (
	"context"
	"pearshop_backend/app/domain/entity"
	"pearshop_backend/app/domain/repository"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) repository.Product {
	return &productRepository{
		db: db,
	}
}

func (repo *productRepository) Create(ctx context.Context, obj *entity.Product) error {

	return nil
}

func (repo *productRepository) Update(ctx context.Context, obj *entity.Product) error {

	return nil
}

func (repo *productRepository) Get(ctx context.Context, spec repository.ISpecs) (entity.Product, error) {

	return entity.Product{}, nil
}

func (repo *productRepository) Find(ctx context.Context, spec repository.ISpecs, paging entity.Paging) ([]entity.Product, error) {

	return []entity.Product{}, nil
}
