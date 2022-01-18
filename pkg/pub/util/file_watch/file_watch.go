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
 * long.qian 2021-11-10 17:05 创建
 */

/**
 * @author long.qian
 */

package file_watch

import (
	"github.com/fsnotify/fsnotify"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/logger"
	"github.com/toolkits/file"
	"io/ioutil"
	"path"
	"runtime"
	"sync"
)

// FileOpEvent 文件事件类型
type FileOpEvent uint32

const (
	// Created 文件创建完毕（已 write 完成）
	Created FileOpEvent = 1 << iota
	// Write 写入事件
	Write
	// Remove 删除事件
	Remove

	Rename
	Chmod
)

type fileEvent struct {
	FilePath string
	Op       FileOpEvent
}

// FileWatch 文件监听
type FileWatch struct {
	// 非 Linux 下的文件创建缓存
	newFileCacheForNonLinux *sync.Map
	// 文件处理channel
	fileHandlerChannel chan *fileEvent
	// 监听器
	watcher *fsnotify.Watcher
	// 判断新文件是否写完的时间更新阈值，默认 500（根据网络环境，自行调整），单位毫秒
	WriteTime int64
	// 是否允许 Debug 日志输出，默认 false
	EnableDebugLog bool
	// 是否处理文件创建事件（Linux 下会利用 in_close_write 事件，性能更好）， 默认 false
	EnableFileCreateHandler bool
	// 是否处理文件删除事件，默认 false
	EnableFileDelHandler bool
	// 是否处理文件的写入事件，默认 false
	EnableFileWriteHandler bool
}

// StartFileWatch 开始目录监听
func (receiver *FileWatch) StartFileWatch(fileHandler func(filePath string, op FileOpEvent), watchDirs ...string) {
	if receiver.WriteTime == 0 {
		receiver.WriteTime = 500
	}
	receiver.fileHandlerChannel = make(chan *fileEvent, runtime.NumCPU()*2)
	go func() {
		// 文件事件
		go func() {
			for {
				e := <-receiver.fileHandlerChannel
				go fileHandler(e.FilePath, e.Op)
			}
		}()
		receiver.eventHandler(watchDirs)
	}()
}

// fsnotify 库的添加监听目录
func (receiver *FileWatch) fsnotifyListening(adsPath string, listeningCount *int32) {
	if file.IsExist(adsPath) {
		if !file.IsFile(adsPath) {
			//监听path目录
			err := receiver.watcher.Add(adsPath)
			if err != nil {
				logger.Sugar.Error("目录：", adsPath, "，监听失败", err)
			} else {
				logger.Sugar.Info("成功监听目录：", adsPath)
				*listeningCount++
			}
			//遍历path目录，监听子目录
			dirs, _ := ioutil.ReadDir(adsPath)
			for _, d := range dirs {
				if d.IsDir() {
					receiver.fsnotifyListening(path.Join(adsPath, d.Name()), listeningCount)
				}
			}
		}
	}
}
