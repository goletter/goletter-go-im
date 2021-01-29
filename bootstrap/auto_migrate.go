package bootstrap

import (
	"goletter-go-im/pkg/database"
	"goletter-go-im/app/models"
)

var MigrateStruct map[string]interface{}

// 初始化表结构体
func init(){
	MigrateStruct = make(map[string]interface{})
	MigrateStruct["user"] = models.User{}
	MigrateStruct["visitor"] = models.Visitor{}
	MigrateStruct["message"] = models.Message{}
	MigrateStruct["welcome"] = models.Welcome{}
}

func autoMigrate()  {
	database.SetMysqlDB()
	for _,v := range MigrateStruct{
		_ = database.DB.AutoMigrate(v)
	}
}
