package handler

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mockUsecase "pearshop_backend/app/usecase/mock"
	mockHashID "pearshop_backend/pkg/hashid/mock"
)

func TestNewProductHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	productFindUsecase := mockUsecase.NewMockProductFind(mockCtrl)
	productCreateUsecase := mockUsecase.NewMockProductCreate(mockCtrl)
	productUpdateUsecase := mockUsecase.NewMockProductUpdate(mockCtrl)
	idHasher := mockHashID.NewMockIDHasher(mockCtrl)

	want := &ProductHandler{
		productFindUsecase:   productFindUsecase,
		productCreateUsecase: productCreateUsecase,
		productUpdateUsecase: productUpdateUsecase,
		idHasher:             idHasher,
	}
	got := NewProductHandler(
		productFindUsecase,
		productUpdateUsecase,
		productCreateUsecase,
		idHasher,
	)

	assert.Equal(t, want, got)
}
