package entity

import "time"

type Product struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Description string
	Price       int

	CreatedAt time.Time `gorm:"->"`
	CreatedBy int

	UpdatedAt time.Time `gorm:"->"`
	UpdatedBy int

	DeletedAt time.Time
}
