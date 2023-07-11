package usecase

import (
	"context"
	"fmt"

	"pearshop_backend/app/domain/entity"
	"pearshop_backend/app/domain/repository"
	"pearshop_backend/app/domain/repository/specifications"
	"pearshop_backend/app/usecase/dto"
)

type productFind struct {
	productRepo repository.Product
}

func NewProductFind(productRepo repository.Product) ProductFind {
	return &productFind{
		productRepo: productRepo,
	}
}

func (uc *productFind) Execute(ctx context.Context, req dto.ProductFindRequest, paging entity.IPaging) ([]entity.Product, error) {
	data, err := uc.productRepo.Find(ctx, specifications.ProductsFind{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}, paging)
	if err != nil {
		return nil, fmt.Errorf("find products: %w", err)
	}

	return data, nil
}
