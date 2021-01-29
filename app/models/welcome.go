package models

import "time"

type Welcome struct {
	ID            uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
	UserId        string `gorm:"type:varchar(100)" json:"user_id"`
	Content       string `gorm:"type:varchar(255)" json:"content"`
	CreatedAt     time.Time `gorm:"column:created_at;index" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}
