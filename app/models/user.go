package models

import (
	"goletter-go-im/pkg/database"
	"time"
)

type User struct {
	ID           uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"kefu_id"`
	Username     string `gorm:"type:varchar(255);not null;unique" json:"username"`
	Nickname     string `gorm:"type:varchar(255)" json:"nickname"`
	Email        string `gorm:"type:varchar(100)" json:"email"`
	Avator       string `gorm:"type:varchar(150)" json:"avator"`
	Password     string `gorm:"type:varchar(150)" json:"password"`
	CreatedAt    time.Time `gorm:"column:created_at;index" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;index" json:"updated_at"`
}

func FindUser(username string) User {
	var user User
	database.DB.Where("username = ?", username).First(&user)
	return user
}