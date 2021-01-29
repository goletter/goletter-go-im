package utils

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net"
	"time"
)

func SetStrToTime(t string) time.Time {
	layout := "2006-01-02 15:04:05" //时间常量
	loc, _ := time.LoadLocation("Asia/Shanghai")
	if len(t) == 10{
		t = fmt.Sprintf("%s 00:00:00",t)
	}
	getTime, _ := time.ParseInLocation(layout, t, loc)
	return getTime
}

// 获取服务器Ip
func GetServerIp() (ip string) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
			}
		}
	}
	return
}

func Uuid() string {
	u2 := uuid.NewV4()
	return u2.String()
}
