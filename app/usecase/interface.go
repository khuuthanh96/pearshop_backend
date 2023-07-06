package usecase

import (
	"context"
	"pearshop_backend/app/domain/entity"
)

type ProductFind interface {
	Execute(ctx context.Context) ([]entity.Product, error)
}

type ProductGet interface {
	Execute(ctx context.Context, id int) (entity.Product, error)
}
