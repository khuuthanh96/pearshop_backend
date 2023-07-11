package usecase

import (
	"context"
	"fmt"

	"pearshop_backend/app/domain/entity"
	"pearshop_backend/app/domain/repository"
	"pearshop_backend/app/usecase/dto"
)

type productCreate struct {
	productRepo repository.Product
}

func NewProductCreate(productRepo repository.Product) ProductCreate {
	return &productCreate{
		productRepo: productRepo,
	}
}

func (uc *productCreate) Execute(ctx context.Context, userID int, req dto.ProductSaveRequest) (entity.Product, error) {
	product := entity.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CreatedBy:   userID,
		UpdatedBy:   userID,
	}

	if err := uc.productRepo.Create(ctx, &product); err != nil {
		return entity.Product{}, fmt.Errorf("create product: %w", err)
	}

	return product, nil
}
