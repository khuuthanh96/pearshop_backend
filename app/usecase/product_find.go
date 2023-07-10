package usecase

import (
	"context"
	"pearshop_backend/app/domain/entity"
	"pearshop_backend/app/domain/repository"
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

func (uc *productFind) Execute(ctx context.Context, req dto.ProductFindRequest, paging entity.Paging) ([]entity.Product, error) {
	return nil, nil
}
