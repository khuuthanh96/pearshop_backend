package payload

import "pearshop_backend/app/domain/entity"

const defaultPageSize uint32 = 30

// PagingRequest holds paging query request
type PagingRequest struct {
	Size *int `form:"size" validate:"omitempty,gte=1,lte=100"` // number of items per page
	Page *int `form:"page" validate:"omitempty,gte=1"`         // page number
}

// StructName returns payload name
func (p PagingRequest) StructName() string {
	return "PagingRequest"
}

// Form converts paging data from request to entity
func (p PagingRequest) Form() entity.PagingRequest {
	pg := entity.PagingRequest{
		Size: defaultPageSize,
		Page: 1,
	}
	if p.Size != nil {
		pg.Size = uint32(*p.Size)
	}

	if p.Page != nil {
		pg.Page = uint32(*p.Page)
	}

	return pg
}

func (p PagingRequest) Validate() error {
	return validate(p)
}
