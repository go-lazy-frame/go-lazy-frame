package go_lazy_frame_example

import (
	"github.com/go-lazy-frame/go-lazy-frame/configs"
)

func initDevConf() {
	Config = SysConfig{
		Config: configs.Config{
			AppName:           "go_lazy_frame_example",
			LogLevel:          "info",
			LogTarget:         "console",
			ApiPrefix:         "/api/v1",
			EnableRbacAuth:    true,
			EnableMon:         true,
			EnableDoc:         true,
			MonPort:           ":10700",
			WebPort:           ":8890",
			MysqlMaxOpenConns: 100,
			MysqlMaxIdleConns: 2,
			MysqlConn:         "admin:admin123@tcp(192.168.0.121:3306)/go_lazy_frame_example?charset=utf8&parseTime=true&loc=Local",
			RedisIP:           "192.168.0.121:6379",
			RedisPassword:     "",
		},
	}
}
