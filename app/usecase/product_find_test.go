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
	"pearshop_backend/app/usecase/dto"
)

func TestNewProductFind(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	productRepo := mockRepo.NewMockProduct(mockCtrl)

	want := &productFind{
		productRepo: productRepo,
	}
	got := NewProductFind(productRepo)

	assert.Equal(t, want, got)
}

func Test_productFind_Execute(t *testing.T) {
	t.Parallel()

	t.Run("find products got error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		productRepo := mockRepo.NewMockProduct(mockCtrl)

		uc := &productFind{
			productRepo: productRepo,
		}
		ctx := context.Background()
		req := dto.ProductFindRequest{}
		paging := entity.NoopPagingRequest{}
		err := errors.New("unexpected error")
		wantErr := fmt.Errorf("find products: %w", err)

		productRepo.EXPECT().Find(ctx, specifications.ProductsFind{
			Name:        req.Name,
			Description: req.Description,
			Price:       req.Price,
		}, paging).Return(nil, err)

		got, gotErr := uc.Execute(ctx, req, paging)

		assert.Nil(t, got)
		assert.Equal(t, wantErr, gotErr)
	})

	t.Run("ok", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		productRepo := mockRepo.NewMockProduct(mockCtrl)

		uc := &productFind{
			productRepo: productRepo,
		}
		name := "test-name"
		ctx := context.Background()
		req := dto.ProductFindRequest{
			Name: &name,
		}
		paging := entity.NoopPagingRequest{}
		want := []entity.Product{
			{
				ID:          1,
				Name:        "aaa-test-1",
				Description: "description-1",
			},
			{
				ID:          2,
				Name:        "bbb-test-1",
				Description: "description-2",
			},
		}

		productRepo.EXPECT().Find(ctx, specifications.ProductsFind{
			Name:        req.Name,
			Description: req.Description,
			Price:       req.Price,
		}, paging).Return([]entity.Product{
			{
				ID:          1,
				Name:        "aaa-test-1",
				Description: "description-1",
			},
			{
				ID:          2,
				Name:        "bbb-test-1",
				Description: "description-2",
			},
		}, nil)

		got, gotErr := uc.Execute(ctx, req, paging)

		assert.Nil(t, gotErr)
		assert.Equal(t, want, got)
	})
}
