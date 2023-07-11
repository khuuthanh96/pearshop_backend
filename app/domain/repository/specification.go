package repository

import "gorm.io/gorm"

type ISpecs interface {
	GormQuery(db *gorm.DB)
}
