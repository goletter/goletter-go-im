package models

import (
	"goletter-go-im/pkg/database"
	"time"
)

type Visitor struct {
	ID           uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
	Name         string `gorm:"type:varchar(255);comment:'昵称'" json:"name"`
	Avator       string `gorm:"type:varchar(150);comment:'头像'" json:"avator"`
	City         string `gorm:"type:varchar(100);comment:'城市'" json:"city"`
	ToId         string `gorm:"type:varchar(100);comment:'对应'" json:"to_id"`
	VisitorId    string `gorm:"type:varchar(100);comment:'游客标识'" json:"visitor_id"`
	Status       uint   `json:"status;comment:'状态'"`
	SourceIp     string `gorm:"type:varchar(100);comment:'客服端Ip'" json:"source_ip"`
	ClientIp     string `gorm:"type:varchar(100);comment:'服务端Ip'" json:"client_ip"`
	CreatedAt    time.Time `gorm:"column:created_at;index" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;index" json:"updated_at"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_at" json:"deleted_at"`
}

func CreateVisitor(name string, avator string, sourceIp string, toId string, visitorId string, city string, clientIp string) {
	old := FindVisitorByVistorId(visitorId)
	if old.Name != "" {
		//更新状态上线
		UpdateVisitor(visitorId, 1, clientIp, sourceIp)
		return
	}

	v := &Visitor{
		Name:      name,
		Avator:    avator,
		SourceIp:  sourceIp,
		ToId:      toId,
		VisitorId: visitorId,
		Status:    1,
		City:      city,
		ClientIp:  clientIp,
	}
	database.DB.Create(v)
}

func FindVisitorByVistorId(visitorId string) Visitor {
	var v Visitor
	database.DB.Where("visitor_id = ?", visitorId).First(&v)
	return v
}

func FindVisitorsOnline() []Visitor {
	var visitors []Visitor
	database.DB.Where("status = ?", 1).Find(&visitors)
	return visitors
}

func UpdateVisitorStatus(visitorId string, status uint) {
	visitor := Visitor{}
	database.DB.Model(&visitor).Where("visitor_id = ?", visitorId).Update("status", status)
}

func UpdateVisitor(visitorId string, status uint, clientIp string, sourceIp string) {
	visitor := &Visitor{
		Status:   status,
		ClientIp: clientIp,
		SourceIp: sourceIp,
	}
	database.DB.Model(&visitor).Where("visitor_id = ?", visitorId).Updates(visitor)
}