package go_lazy_frame_example

import (
	"fmt"
	"github.com/go-lazy-frame/go-lazy-frame/configs"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/logger"
)

var (
	// Config 系统配置
	Config SysConfig
)

// SysConfig 系统配置项
type SysConfig struct {
	configs.Config
}

func init() {
	ConfInit()
}

// ConfInit 初始化配置
func ConfInit() {
	env := configs.GetEnv()
	switch env {
	case configs.DevEnv:
		fmt.Println("=============================================================================")
		fmt.Println("================================== 开发环境模式 ===============================")
		fmt.Println("=============================================================================")
		initDevConf()
	case configs.ProdEnv:
		fmt.Println("=============================================================================")
		fmt.Println("================================== 线上环境模式 ===============================")
		fmt.Println("=============================================================================")
		initProdConf()
	default:
		fmt.Println("=============================================================================")
		fmt.Println("================================== 本地环境模式 ===============================")
		fmt.Println("=============================================================================")
		initLocalConf()
	}
	configs.GeneralConfig = &Config.Config
	// 从环境变量更新配置
	configs.ConfigByEnv(&Config)

	// 加载日志配置
	logger.Level = configs.GetLoggerLevelByConfig(Config.LogLevel)
	logger.Target = configs.GetLoggerTargetByConfig(Config.LogTarget)

	// 初始化日志
	logger.Init(configs.GeneralConfig.AppName, Config.LogMaxSize, Config.LogMaxBackups, Config.LogMaxAge)
}
