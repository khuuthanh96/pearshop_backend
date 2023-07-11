package payload

import "pearshop_backend/app/usecase/dto"

type ProductFindRequest struct {
	Name        *string  `form:"name" validate:"omitempty,max=255"`
	Description *string  `form:"description" validate:"omitempty,max=1000"`
	Price       *float64 `form:"price"`
	PagingRequest
}

func (p ProductFindRequest) StructName() string {
	return "ProductFindRequest"
}

func (p ProductFindRequest) Validate() error {
	if err := validate(p); err != nil {
		return err
	}

	return validate(p.PagingRequest)
}

func (p ProductFindRequest) ToDTO() dto.ProductFindRequest {
	return dto.ProductFindRequest{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}
}

type ProductSaveRequest struct {
	Name        string  `json:"name" validate:"required,max=255"`
	Description string  `json:"description" validate:"omitempty,max=1000"`
	Price       float64 `json:"price" validate:"omitempty,gte=0"`
}

func (p ProductSaveRequest) StructName() string {
	return "ProductSaveRequest"
}

func (p ProductSaveRequest) Validate() error {
	return validate(p)
}

func (p ProductSaveRequest) ToDTO() dto.ProductSaveRequest {
	return dto.ProductSaveRequest{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}
}
