package go_lazy_frame_example

import (
	"github.com/go-lazy-frame/go-lazy-frame/configs"
)

func initProdConf() {
	Config = SysConfig{
		Config: configs.Config{
			AppName:           "go_lazy_frame_example",
			LogLevel:          "info",
			LogTarget:         "file",
			ApiPrefix:         "/api/v1",
			EnableRbacAuth:    true,
			EnableMon:         false,
			EnableDoc:         true,
			MonPort:           ":10700",
			WebPort:           ":8890",
			MysqlMaxOpenConns: 100,
			MysqlMaxIdleConns: 2,
			MysqlConn:         "user:password@tcp(127.0.0.1:3306)/go_lazy_frame_example?charset=utf8&parseTime=true&loc=Local",
			RedisIP:           "127.0.0.1:6379",
			RedisPassword:     "",
		},
	}
}
