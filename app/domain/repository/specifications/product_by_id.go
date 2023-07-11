package specifications

import (
	"gorm.io/gorm"

	"pearshop_backend/app/domain/entity"
	"pearshop_backend/app/domain/repository"
)

type productByID struct {
	id int
}

func ProductByID(id int) repository.ISpecs {
	return &productByID{
		id: id,
	}
}

func (spec *productByID) GormQuery(db *gorm.DB) {
	db = db.Model(&entity.Product{}).Where("id = ?", spec.id)
}
