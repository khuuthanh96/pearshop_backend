package usecase

import (
	"context"
	"fmt"

	"pearshop_backend/app/domain/entity"
	"pearshop_backend/app/domain/repository"
	"pearshop_backend/app/domain/repository/specifications"
	appErrors "pearshop_backend/app/errors"
	"pearshop_backend/app/usecase/dto"
	"pearshop_backend/pkg/hashid"
)

type productUpdate struct {
	productRepo repository.Product
	idHasher    hashid.IDHasher
}

func NewProductUpdate(productRepo repository.Product, idHasher hashid.IDHasher) ProductUpdate {
	return &productUpdate{
		productRepo: productRepo,
		idHasher:    idHasher,
	}
}

func (uc *productUpdate) Execute(ctx context.Context, userID, productID int, req dto.ProductSaveRequest) (entity.Product, error) {
	product, err := uc.productRepo.Get(ctx, specifications.ProductByID(productID))
	if err != nil {
		if err == appErrors.ErrRecordNotFound {
			return entity.Product{}, appErrors.NewNotFoundErr(
				appErrors.CodeProductNotFound,
				"product not found",
				uc.idHasher.Encode(productID),
			)
		}

		return entity.Product{}, fmt.Errorf("find product by id: %w", err)
	}

	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.UpdatedBy = userID

	if err := uc.productRepo.Update(ctx, &product); err != nil {
		return entity.Product{}, fmt.Errorf("update product: %w", err)
	}

	return product, nil
}
