package usecase

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"pearshop_backend/app/domain/entity"
	mockRepo "pearshop_backend/app/domain/repository/mock"
	"pearshop_backend/app/domain/repository/specifications"
	appErrors "pearshop_backend/app/errors"
	"pearshop_backend/app/usecase/dto"
	mockHashID "pearshop_backend/pkg/hashid/mock"
)

func TestNewProductUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	productRepo := mockRepo.NewMockProduct(mockCtrl)
	idHasher := mockHashID.NewMockIDHasher(mockCtrl)

	want := &productUpdate{
		productRepo: productRepo,
		idHasher:    idHasher,
	}
	got := NewProductUpdate(productRepo, idHasher)

	assert.Equal(t, want, got)
}

func Test_productUpdate_Execute(t *testing.T) {
	t.Parallel()

	t.Run("get product by id got error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		productRepo := mockRepo.NewMockProduct(mockCtrl)
		idHasher := mockHashID.NewMockIDHasher(mockCtrl)

		uc := &productUpdate{
			productRepo: productRepo,
			idHasher:    idHasher,
		}
		ctx := context.Background()
		req := dto.ProductSaveRequest{}
		userID := 1
		productID := 2
		err := errors.New("unexpected error")
		wantErr := fmt.Errorf("find product by id: %w", err)
		want := entity.Product{}

		productRepo.EXPECT().Get(ctx, specifications.ProductByID(productID)).Return(entity.Product{}, err)

		got, gotErr := uc.Execute(ctx, userID, productID, req)

		assert.Equal(t, want, got)
		assert.Equal(t, wantErr, gotErr)
	})

	t.Run("get product by id not found", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		productRepo := mockRepo.NewMockProduct(mockCtrl)
		idHasher := mockHashID.NewMockIDHasher(mockCtrl)

		uc := &productUpdate{
			productRepo: productRepo,
			idHasher:    idHasher,
		}
		ctx := context.Background()
		req := dto.ProductSaveRequest{}
		userID := 1
		productID := 2
		encodedProductID := "encoded-product-id"
		err := appErrors.ErrRecordNotFound
		wantErr := appErrors.NewNotFoundErr(
			appErrors.CodeProductNotFound,
			"product not found",
			encodedProductID,
		)
		want := entity.Product{}
		productRepo.EXPECT().Get(ctx, specifications.ProductByID(productID)).Return(entity.Product{}, err)
		idHasher.EXPECT().Encode(productID).Return(encodedProductID)

		got, gotErr := uc.Execute(ctx, userID, productID, req)

		assert.Equal(t, want, got)
		assert.Equal(t, wantErr, gotErr)
	})

	t.Run("update product got error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		productRepo := mockRepo.NewMockProduct(mockCtrl)
		idHasher := mockHashID.NewMockIDHasher(mockCtrl)

		uc := &productUpdate{
			productRepo: productRepo,
			idHasher:    idHasher,
		}
		ctx := context.Background()
		req := dto.ProductSaveRequest{
			Name:        "name-2",
			Description: "description-2",
		}
		userID := 1
		productID := 2
		err := errors.New("unexpected error")
		wantErr := fmt.Errorf("update product: %w", err)

		product := entity.Product{
			ID:          productID,
			Name:        "name-1",
			Description: "description-1",
			CreatedBy:   userID,
			UpdatedBy:   userID,
		}
		want := entity.Product{}

		productRepo.EXPECT().Get(ctx, specifications.ProductByID(productID)).Return(product, nil)
		productRepo.EXPECT().Update(ctx, &entity.Product{
			ID:          productID,
			Name:        "name-2",
			Description: "description-2",
			CreatedBy:   userID,
			UpdatedBy:   userID,
		}).Return(err)

		got, gotErr := uc.Execute(ctx, userID, productID, req)

		assert.Equal(t, want, got)
		assert.Equal(t, wantErr, gotErr)
	})

	t.Run("ok", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		productRepo := mockRepo.NewMockProduct(mockCtrl)

		uc := &productUpdate{
			productRepo: productRepo,
		}
		ctx := context.Background()
		req := dto.ProductSaveRequest{
			Name:        "name-2",
			Description: "description-2",
		}
		userID := 1
		productID := 2
		product := entity.Product{
			ID:          productID,
			Name:        "name-1",
			Description: "description-1",
			CreatedBy:   userID,
			UpdatedBy:   userID,
		}
		want := entity.Product{
			ID:          productID,
			Name:        "name-2",
			Description: "description-2",
			CreatedBy:   userID,
			UpdatedBy:   userID,
		}

		productRepo.EXPECT().Get(ctx, specifications.ProductByID(productID)).Return(product, nil)
		productRepo.EXPECT().Update(ctx, &entity.Product{
			ID:          productID,
			Name:        "name-2",
			Description: "description-2",
			CreatedBy:   userID,
			UpdatedBy:   userID,
		}).Return(nil)

		got, gotErr := uc.Execute(ctx, userID, productID, req)

		assert.Nil(t, gotErr)
		assert.Equal(t, want, got)
	})
}
