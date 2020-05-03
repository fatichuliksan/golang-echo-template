package model

import "time"

// User Attribute
type User struct {
	DefaultAttribute
	Name            string `gorm:"type:varchar(191)"`
	Email           string `gorm:"type:varchar(191)"`
	EmailVerifiedAt *time.Time
	Password        string `gorm:"type:varchar(191)"`
	Avatar          string
	RememberToken   string `gorm:"type:varchar(100)"`
}

// TableName ...
func (User) TableName() string {
	return "users"
}
