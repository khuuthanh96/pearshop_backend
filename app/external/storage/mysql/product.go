package mysql

import (
	"context"

	"gorm.io/gorm"

	"pearshop_backend/app/domain/entity"
	"pearshop_backend/app/domain/repository"
	appErrors "pearshop_backend/app/errors"
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
	return repo.db.WithContext(ctx).Create(&obj).Error
}

func (repo *productRepository) Update(ctx context.Context, obj *entity.Product) error {
	return repo.db.WithContext(ctx).Updates(&obj).Error
}

func (repo *productRepository) Get(ctx context.Context, spec repository.ISpecs) (entity.Product, error) {
	res := entity.Product{}
	db := repo.db.WithContext(ctx)

	spec.GormQuery(db)

	if err := db.Take(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return res, appErrors.ErrRecordNotFound
		}

		return res, err
	}

	return res, nil
}

func (repo *productRepository) Find(ctx context.Context, spec repository.ISpecs, paging entity.IPaging) ([]entity.Product, error) {
	res := []entity.Product{}
	db := repo.db.WithContext(ctx)

	spec.GormQuery(db)
	paging.GormPaging(db)

	if err := db.Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (repo *productRepository) AssignTx(txm repository.TransactionManager) {
	if v, ok := txm.GetTx().(*gorm.DB); ok && v != nil {
		repo.db = v
	}
}
