package model

import "time"

// DefaultAttribute definition
type DefaultAttribute struct {
	ID        uint       `gorm:"primary_key;column:id"`
	CreatedBy *uint      `gorm:"column:created_by"`
	UpdatedBy *uint      `gorm:"column:updated_by"`
	DeletedBy *uint      `gorm:"column:deleted_by"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}
