package entity

import (
	"time"
	"gorm.io/gorm"
)

type Books struct {
	Id int `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Author string `json:"author"`
	Cover string `json:"cover"`
	CreatedAt time.Time `json:"created_at",gorm:"autoUpdateTime:nano"`     // Automatically managed by GORM for creation time
	UpdatedAt time.Time `json:"updated_at"` // Automatically managed by GORM for update time
	DeletedAt gorm.DeletedAt `json:"deleted_at; gorm:"index,column:deleted_at"`
}