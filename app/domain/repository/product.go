package repository

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"context"

	"pearshop_backend/app/domain/entity"
)

type Product interface {
	Create(ctx context.Context, obj *entity.Product) error
	Update(ctx context.Context, obj *entity.Product) error
	Get(ctx context.Context, spec ISpecs) (entity.Product, error)
	Find(ctx context.Context, spec ISpecs, paging entity.IPaging) ([]entity.Product, error)
	AssignTx(tx TransactionManager)
}
