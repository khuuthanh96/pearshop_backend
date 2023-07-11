package specifications

import (
	"gorm.io/gorm"

	"pearshop_backend/app/domain/entity"
)

type ProductsFind struct {
	Name        *string
	Description *string
	Price       *float64
}

func (spec ProductsFind) GormQuery(db *gorm.DB) {
	db = db.Model(&entity.Product{})

	if spec.Name != nil {
		db = db.Where("name like %%%s%", spec.Name)
	}

	if spec.Description != nil {
		db = db.Where("description like %%%s%", spec.Description)
	}

	if spec.Price != nil {
		db = db.Where("price = ?", spec.Price)
	}
}
