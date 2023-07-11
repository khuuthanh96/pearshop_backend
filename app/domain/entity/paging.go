package entity

import "gorm.io/gorm"

type IPaging interface {
	GormPaging(db *gorm.DB)
}

// NoopPagingRequest for no pagination
type NoopPagingRequest struct{}

func (p NoopPagingRequest) GormPaging(db *gorm.DB) {
	return
}

// PagingRequest holds paging information
type PagingRequest struct {
	Size uint32
	Page uint32
}

func (p PagingRequest) GormPaging(db *gorm.DB) {
	db.Limit(int(p.Size)).Offset(int(p.Size * (p.Page - 1)))
}
