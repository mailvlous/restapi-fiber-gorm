package entity

import (
	"time"
	"gorm.io/gorm"
)

type Users struct {
	Id int `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password  string `json:"-" gorm:"column:password"`
	Phone string `json:"phone"`
	Role string `json:"role"`
	CreatedAt time.Time `json:"created_at",gorm:"autoUpdateTime:nano"`     // Automatically managed by GORM for creation time
	UpdatedAt time.Time `json:"updated_at"` // Automatically managed by GORM for update time
	DeletedAt gorm.DeletedAt `json:"deleted_at; gorm:"index,column:deleted_at"`
}