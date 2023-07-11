package dto

type ProductFindRequest struct {
	Name        *string
	Description *string
	Price       *float64
}

type ProductSaveRequest struct {
	Name        string
	Description string
	Price       float64
}
