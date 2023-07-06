package entity

import "time"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       int

	CreatedAt time.Time
	CreatedBy int

	UpdatedAt time.Time
	UpdatedBy int

	DeletedAt time.Time
}
