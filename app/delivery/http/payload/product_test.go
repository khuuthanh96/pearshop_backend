package payload

import (
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	appErrors "pearshop_backend/app/errors"
	"pearshop_backend/app/usecase/dto"
)

func TestProductFindRequest_StructName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "success",
			want: "ProductFindRequest",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payl := ProductFindRequest{}
			if got := payl.StructName(); got != tt.want {
				t.Errorf("ProductFindRequest.StructName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductFindRequest_Validate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	tests := []struct {
		name    string
		fields  ProductFindRequest
		wantErr error
	}{
		{
			name: "violate `max` tag",
			fields: ProductFindRequest{
				Name: func() *string {
					s := strings.Repeat("a", 256)

					return &s
				}(),
				Description: func() *string {
					s := strings.Repeat("a", 1001)

					return &s
				}(),
			},
			wantErr: appErrors.SystemErrors{
				appErrors.NewInvalidArgumentErr(
					"CODE_PRODUCT_FIND_INVALID_NAME",
					"name must not exceed 255 characters",
					strings.Repeat("a", 256),
				),
				appErrors.NewInvalidArgumentErr(
					"CODE_PRODUCT_FIND_INVALID_DESCRIPTION",
					"description must not exceed 1000 characters",
					strings.Repeat("a", 1001),
				),
			},
		},
		{
			name: "success",
			fields: ProductFindRequest{
				Name: func() *string {
					s := strings.Repeat("a", 255)

					return &s
				}(),
				Description: func() *string {
					s := strings.Repeat("a", 1000)

					return &s
				}(),
				PagingRequest: PagingRequest{
					Size: func() *int {
						i := 1

						return &i
					}(),
					Page: func() *int {
						i := 30

						return &i
					}(),
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payl := ProductFindRequest{
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				Price:       tt.fields.Price,
				PagingRequest: PagingRequest{
					Size: tt.fields.PagingRequest.Size,
					Page: tt.fields.PagingRequest.Page,
				},
			}
			got := payl.Validate()

			assert.Equal(t, tt.wantErr, got)
		})
	}
}

func TestCreateAnnouncementRequest_ToDTO(t *testing.T) {
	tests := []struct {
		name   string
		fields ProductFindRequest
		want   dto.ProductFindRequest
	}{
		{
			name: "success",
			fields: ProductFindRequest{
				Name: func() *string {
					s := "name"

					return &s
				}(),
				Description: func() *string {
					s := "description"

					return &s
				}(),
				Price: func() *float64 {
					n := float64(10)

					return &n
				}(),
			},
			want: dto.ProductFindRequest{
				Name: func() *string {
					s := "name"

					return &s
				}(),
				Description: func() *string {
					s := "description"

					return &s
				}(),
				Price: func() *float64 {
					n := float64(10)

					return &n
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payl := ProductFindRequest{
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				Price:       tt.fields.Price,
			}
			got := payl.ToDTO()

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestProductSaveRequest_StructName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "success",
			want: "ProductSaveRequest",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payl := ProductSaveRequest{}
			if got := payl.StructName(); got != tt.want {
				t.Errorf("ProductSaveRequest.StructName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductSaveRequest_Validate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	tests := []struct {
		name    string
		fields  ProductSaveRequest
		wantErr error
	}{
		{
			name:   "violate `max` tag",
			fields: ProductSaveRequest{},
			wantErr: appErrors.SystemErrors{
				appErrors.NewInvalidArgumentErr(
					"CODE_PRODUCT_INVALID_NAME",
					"name is required",
					"",
				),
			},
		},
		{
			name: "violate `max` tag",
			fields: ProductSaveRequest{
				Name:        strings.Repeat("a", 256),
				Description: strings.Repeat("a", 1001),
			},
			wantErr: appErrors.SystemErrors{
				appErrors.NewInvalidArgumentErr(
					"CODE_PRODUCT_INVALID_NAME",
					"name must not exceed 255 characters",
					strings.Repeat("a", 256),
				),
				appErrors.NewInvalidArgumentErr(
					"CODE_PRODUCT_INVALID_DESCRIPTION",
					"description must not exceed 1000 characters",
					strings.Repeat("a", 1001),
				),
			},
		},
		{
			name: "violate `gte` tag",
			fields: ProductSaveRequest{
				Name:        strings.Repeat("a", 255),
				Description: strings.Repeat("a", 1000),
				Price:       -1,
			},
			wantErr: appErrors.SystemErrors{
				appErrors.NewInvalidArgumentErr(
					"CODE_PRODUCT_INVALID_PRICE",
					"price must equal or greater than 0",
					float64(-1),
				),
			},
		},
		{
			name: "success",
			fields: ProductSaveRequest{
				Name:        strings.Repeat("a", 255),
				Description: strings.Repeat("a", 1000),
				Price:       10,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payl := ProductSaveRequest{
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				Price:       tt.fields.Price,
			}
			got := payl.Validate()

			assert.Equal(t, tt.wantErr, got)
		})
	}
}

func TestProductSaveRequest_ToDTO(t *testing.T) {
	tests := []struct {
		name   string
		fields ProductSaveRequest
		want   dto.ProductSaveRequest
	}{
		{
			name: "success",
			fields: ProductSaveRequest{
				Name:        strings.Repeat("a", 255),
				Description: strings.Repeat("a", 1000),
				Price:       10,
			},
			want: dto.ProductSaveRequest{
				Name:        strings.Repeat("a", 255),
				Description: strings.Repeat("a", 1000),
				Price:       10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payl := ProductSaveRequest{
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				Price:       tt.fields.Price,
			}
			got := payl.ToDTO()

			assert.Equal(t, tt.want, got)
		})
	}
}
