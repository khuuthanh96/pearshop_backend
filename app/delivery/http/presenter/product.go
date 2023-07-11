package presenter

import (
	"time"

	"pearshop_backend/app/domain/entity"
	"pearshop_backend/pkg/hashid"
)

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FormProduct(idHasher hashid.IDHasher, e entity.Product) Product {
	return Product{
		ID:          idHasher.Encode(e.ID),
		Name:        e.Name,
		Description: e.Description,
		Price:       e.Price,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}

func FormProducts(idHasher hashid.IDHasher, es []entity.Product) []Product {
	res := make([]Product, len(es))

	for i := range es {
		res[i] = FormProduct(idHasher, es[i])
	}

	return res
}
