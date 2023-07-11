package usecase

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"context"

	"pearshop_backend/app/domain/entity"
	"pearshop_backend/app/usecase/dto"
)

type ProductFind interface {
	Execute(ctx context.Context, req dto.ProductFindRequest, paging entity.IPaging) ([]entity.Product, error)
}

type ProductUpdate interface {
	Execute(ctx context.Context, userID, productID int, req dto.ProductSaveRequest) (entity.Product, error)
}

type ProductCreate interface {
	Execute(ctx context.Context, userID int, req dto.ProductSaveRequest) (entity.Product, error)
}
