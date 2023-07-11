package presenter

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"pearshop_backend/app/domain/entity"
	mockHashID "pearshop_backend/pkg/hashid/mock"
)

func TestFromProduct(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	idHasher := mockHashID.NewMockIDHasher(mockCtrl)

	encodedID := "encoded-id"

	idHasher.EXPECT().Encode(1).Return(encodedID)

	p := entity.Product{
		ID:          1,
		Name:        "product-name",
		Description: "product-description",
		Price:       100,
	}

	want := Product{
		ID:          encodedID,
		Name:        "product-name",
		Description: "product-description",
		Price:       100,
	}
	got := FormProduct(idHasher, p)

	assert.Equal(t, want, got)
}

func TestFromProducts(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	idHasher := mockHashID.NewMockIDHasher(mockCtrl)

	encodedID1 := "encoded-id-1"
	encodedID2 := "encoded-id-2"

	idHasher.EXPECT().Encode(1).Return(encodedID1)
	idHasher.EXPECT().Encode(2).Return(encodedID2)

	ps := []entity.Product{
		{
			ID:          1,
			Name:        "product-name-1",
			Description: "product-description-1",
			Price:       100,
		},
		{
			ID:          2,
			Name:        "product-name-2",
			Description: "product-description-2",
			Price:       200,
		},
	}

	want := []Product{
		{
			ID:          encodedID1,
			Name:        "product-name-1",
			Description: "product-description-1",
			Price:       100,
		},
		{
			ID:          encodedID2,
			Name:        "product-name-2",
			Description: "product-description-2",
			Price:       200,
		},
	}
	got := FormProducts(idHasher, ps)

	assert.Equal(t, want, got)
}
