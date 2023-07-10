package dto

type ProductFindRequest struct {
	Name        *string
	Description *string
	Price       *int
}

type ProductSaveRequest struct {
	Name        string
	Description string
	Price       int
}
