//             ,%%%%%%%%,
//           ,%%/\%%%%/\%%
//          ,%%%\c "" J/%%%
// %.       %%%%/ o  o \%%%
// `%%.     %%%%    _  |%%%
//  `%%     `%%%%(__Y__)%%'
//  //       ;%%%%`\-/%%%'
// ((       /  `%%%%%%%'
//  \\    .'          |
//   \\  /       \  | |
//    \\/攻城狮保佑) | |
//     \         /_ | |__
//     (___________)))))))                   `\/'
/*
 * 修订记录:
 * long.qian 2021-10-02 18:09 创建
 */

/**
 * @author long.qian
 */

package configs

import (
	"fmt"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/logger"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"reflect"
	"strings"
)

const (
	// EnvKey 环境 Key，通过 os.Getenv("RUN_ENV") 获取
	EnvKey = "RUN_ENV"
	// ProdEnv 线上环境标识
	ProdEnv = "prod"
	// DevEnv 开发环境标识
	DevEnv = "dev"
	// LocalEnv 本地环境标识
	LocalEnv = "local"
)

var (
	// GeneralConfig 项目通用配置，由应用启动时注入
	GeneralConfig *Config
)

// Config 项目公共配置
type Config struct {
	AppName           string   `json:"app_name"`             // 应用名
	LogLevel          string   `json:"log_level"`            // 日志级别，只能为：debug,info,warn,error,dpanic,panic,fatal
	LogTarget         string   `json:"log_target"`           // 日志输出目标，只能为：file,console
	EnableRbacAuth    bool     `json:"enable_rbac_auth"`     // 是否启用 RBAC 权限系统，若启用，则自动注入相关接口和数据库数据
	EnableMon         bool     `json:"enable_mon"`           // 是否开启监控
	EnableDoc         bool     `json:"enable_doc"`           // 是否开启接口文档
	MonPort           string   `json:"mon_port"`             // 监控服务监听端口
	WebPort           string   `json:"web_port"`             // Web 服务监听端口
	ApiPrefix         string   `json:"api_prefix"`           // API 接口地址的前缀
	MysqlMaxOpenConns int      `json:"mysql_max_open_conns"` // 0：不限制（注意是否会触发到 mysql 数据库的最大连接数）
	MysqlMaxIdleConns int      `json:"mysql_max_idle_conns"`
	MysqlConn         string   `json:"mysql_conn"`     // 数据库连接地址
	RedisIP           string   `json:"redis_ip"`       // Redis 连接地址
	RedisPassword     string   `json:"redis_password"` // Redis 连接密码
	NotAuthUrl        []string `json:"not_auth_url"`   // 不用鉴权的访问地址前缀
	Test              string   `json:"test"`
}

func GetLoggerLevelByConfig(logLevel string) zapcore.Level {
	fmt.Println("日志级别为：", logLevel)
	switch logLevel {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "dpanic":
		return zap.DPanicLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	default:
		fmt.Println("日志级别只能为：debug,info,warn,error,dpanic,panic,fatal，默认为 info 级别")
		return zap.InfoLevel
	}
}

func GetLoggerTargetByConfig(logTarget string) string {
	fmt.Println("日志输出目标为：", logTarget)
	switch logTarget {
	case "file":
		return logger.File
	case "console":
		return logger.Console
	default:
		fmt.Println("日志输出目标只能为：file,console，默认为 console")
		return logger.Console
	}
}

// GetEnv 获取当前运行环境
func GetEnv() string {
	env := os.Getenv(EnvKey)
	if env == "" {
		env = LocalEnv
	}
	return env
}

// IsLocalEnv 是否是本地环境
func IsLocalEnv() bool {
	return GetEnv() == LocalEnv
}

// IsDevEnv 是否是开发环境
func IsDevEnv() bool {
	return GetEnv() == DevEnv
}

// IsProdEnv 是否是线上环境
func IsProdEnv() bool {
	return GetEnv() == ProdEnv
}

// ConfigByEnv 根据环境变量 ENV_xxx（xxx 为对应的具体配置名），更新配置
// 注意：只更新字段类型为 string,int,int64,int32,float32,float64,bool 的配置，对应的集合类型不会更新
func ConfigByEnv(obj interface{}) {
	fieldType := reflect.TypeOf(obj)
	fieldValue := reflect.ValueOf(obj)
	numField := fieldType.Elem().NumField()
	for i := 0; i < numField; i++ {
		fieldName := fieldType.Elem().Field(i).Name
		typeName := fieldType.Elem().Field(i).Type.Name()
		if typeName == "Config" {
			v := fieldValue.Elem().Field(i)
			ConfigByEnv(v.Addr().Interface())
			continue
		}
		valueEnv := os.Getenv("ENV_" + fieldName)
		if valueEnv != "" {
			switch typeName {
			case "string":
				f := fieldValue.Elem().Field(i)
				if f.CanSet() {
					fieldValue.Elem().Field(i).SetString(valueEnv)
					fmt.Printf("配置 %s 已根据环境变量，更新为：%v\n", fieldName, valueEnv)
				} else {
					fmt.Printf("配置 %s 根据环境变量更新失败：不能设置\n", fieldName)
				}
				break
			case "int":
			case "int32":
			case "int64":
				f := fieldValue.Elem().Field(i)
				if f.CanSet() {
					fieldValue.Elem().Field(i).SetInt(util.NumberUtil.StringToInt64(valueEnv))
					fmt.Printf("配置 %s 已根据环境变量，更新为：%v\n", fieldName, valueEnv)
				} else {
					fmt.Printf("配置 %s 根据环境变量更新失败：不能设置\n", fieldName)
				}
				break
			case "float32":
			case "float64":
				f := fieldValue.Elem().Field(i)
				if f.CanSet() {
					fieldValue.Elem().Field(i).SetFloat(util.NumberUtil.StringToFloat64(valueEnv))
					fmt.Printf("配置 %s 已根据环境变量，更新为：%v\n", fieldName, valueEnv)
				} else {
					fmt.Printf("配置 %s 根据环境变量更新失败：不能设置\n", fieldName)
				}
				break
			case "bool":
				f := fieldValue.Elem().Field(i)
				if f.CanSet() {
					fieldValue.Elem().Field(i).SetBool(strings.ToLower(valueEnv) == "true")
					fmt.Printf("配置 %s 已根据环境变量，更新为：%v\n", fieldName, strings.ToLower(valueEnv) == "true")
				} else {
					fmt.Printf("配置 %s 根据环境变量更新失败：不能设置\n", fieldName)
				}
				break
			default:
				fmt.Printf("配置 %s 根据环境变量更新失败：只能为 string,int,int64,int32,float32,float64,bool 类型\n", fieldName)
			}
		}

	}
}
