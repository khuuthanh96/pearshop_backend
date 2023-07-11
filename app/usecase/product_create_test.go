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
	"pearshop_backend/app/usecase/dto"
)

func TestNewProductCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	productRepo := mockRepo.NewMockProduct(mockCtrl)

	want := &productCreate{
		productRepo: productRepo,
	}
	got := NewProductCreate(productRepo)

	assert.Equal(t, want, got)
}

func Test_productCreate_Execute(t *testing.T) {
	t.Parallel()

	t.Run("create product got error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		productRepo := mockRepo.NewMockProduct(mockCtrl)

		uc := &productCreate{
			productRepo: productRepo,
		}
		ctx := context.Background()
		req := dto.ProductSaveRequest{}
		userID := 1
		err := errors.New("unexpected error")
		wantErr := fmt.Errorf("create product: %w", err)
		want := entity.Product{}

		productRepo.EXPECT().Create(ctx, &entity.Product{
			CreatedBy: 1,
			UpdatedBy: 1,
		}).Return(err)

		got, gotErr := uc.Execute(ctx, userID, req)

		assert.Equal(t, want, got)
		assert.Equal(t, wantErr, gotErr)
	})

	t.Run("ok", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		productRepo := mockRepo.NewMockProduct(mockCtrl)

		uc := &productCreate{
			productRepo: productRepo,
		}
		ctx := context.Background()
		req := dto.ProductSaveRequest{
			Name:        "product",
			Description: "test product",
			Price:       100,
		}
		userID := 1
		want := entity.Product{
			ID:          1,
			Name:        "product",
			Description: "test product",
			Price:       100,
			CreatedBy:   1,
			UpdatedBy:   1,
		}

		productRepo.EXPECT().Create(ctx, &entity.Product{
			Name:        "product",
			Description: "test product",
			Price:       100,
			CreatedBy:   1,
			UpdatedBy:   1,
		}).SetArg(1, want).Return(nil)

		got, gotErr := uc.Execute(ctx, userID, req)

		assert.Nil(t, gotErr)
		assert.Equal(t, want, got)
	})
}
