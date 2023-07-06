package repository

import (
	"context"

	"pearshop_backend/app/domain/entity"
)

type Product interface {
	Create(ctx context.Context, obj *entity.Product) error
	Update(ctx context.Context, obj *entity.Product) error
	Get(ctx context.Context, id int) (entity.Product, error)
	Find(ctx context.Context) ([]entity.Product, error)
}
