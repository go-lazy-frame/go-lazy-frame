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
 * long.qian 2021-10-04 18:32 创建
 */

/**
 * @author long.qian
 */

package util

import (
	"github.com/go-ping/ping"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/logger"
	"net"
	"runtime"
	"strings"
	"time"
)

var (
	IpUtil = new(ipUtil)
)

type ipUtil struct {

}

// GetLocalIp 获取本地IP
func (me *ipUtil) GetLocalIp() string {
	ip := ""
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				if strings.HasPrefix(ip, "192.168") || strings.HasPrefix(ip, "10.") || strings.HasPrefix(ip, "172.") {
					break
				}
			}

		}
	}
	if ip == "" {
		return "127.0.0.1"
	}
	return ip
}

func (me *ipUtil) IsPing(host string, timeout time.Duration) bool {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		logger.Sugar.Error(err)
		return false
	}
	pinger.Count = 1
	pinger.Timeout = timeout
	if runtime.GOOS == "windows" {
		pinger.SetPrivileged(true)
	}
	// Linux 上，注意查看：https://github.com/go-ping/ping/blob/master/README.md#supported-operating-systems
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		logger.Sugar.Error(err)
		if runtime.GOOS == "linux" {
			logger.Sugar.Info(`Linux 平台，您可能需要执行以下命令：sudo sysctl -w net.ipv4.ping_group_range="0   2147483647"`)
		}
		return false
	}
	stats := pinger.Statistics()
	logger.Sugar.Info("Ping Packets Received Count：", stats.PacketsRecv)
	if stats.PacketsRecv >= 1 {
		return true
	}
	return false
}
