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
 * long.qian 2021-11-11 10:05 创建
 */

/**
 * @author long.qian
 */

package file_watch

import (
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/logger"
	"testing"
)

func init() {
	logger.Init("sand_server")
}

func TestFileWatch(t *testing.T) {
	watch := FileWatch{}
	watch.StartFileWatch(func(newFile string) {
		logger.Sugar.Info("新文件：", newFile)
	}, "/Volumes/E-NTFS/test")
	select {}
}

