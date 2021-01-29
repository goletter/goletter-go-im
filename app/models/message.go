package models

import (
	"goletter-go-im/pkg/database"
	"time"
)

type Message struct {
	ID            uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
	KefuId        string `gorm:"type:varchar(100)" json:"kefu_id"`
	VisitorId     string `gorm:"type:varchar(100)" json:"visitor_id"`
	Content       string `gorm:"type:varchar(255)" json:"content"`
	MesType       string `gorm:"type:varchar(100)" json:"mes_type"`
	Status        string `gorm:"type:varchar(50)" json:"status"`
	CreatedAt     time.Time `gorm:"column:created_at;index" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func FindMessageByVisitorId(visitor_id string) []Message {
	var messages []Message
	database.DB.Where("visitor_id=?", visitor_id).Order("id asc").Find(&messages)
	return messages
}

//修改消息状态
func ReadMessageByVisitorId(visitor_id string) {
	database.DB.Table("messages").Where("visitor_id=?", visitor_id).Update("status", "read")
}