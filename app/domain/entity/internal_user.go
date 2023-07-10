package entity

import "time"

type RoleType int8

const (
	RoleTypeNormal RoleType = iota + 1
	RoleTypeManager
	RoleTypeAdmin
)

type InternalUser struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	RoleType  RoleType
	CreatedAt time.Time `gorm:"->"`
	UpdatedAt time.Time `gorm:"->"`
}
