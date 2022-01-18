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
 * long.qian 2022-01-17 15:16 创建
 */

/**
 * @author long.qian
 */

package file_watch

import (
	"github.com/fsnotify/fsnotify"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/logger"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/util"
	"github.com/toolkits/file"
	"os"
	"strings"
	"sync"
	"time"
)

// 非 linux 下的处理
func (receiver *FileWatch) nonLinuxEventHandler(watchDirs []string) {
	logger.Sugar.Info("非 Linux 文件事件处理")
	if receiver.EnableFileCreateHandler {
		receiver.newFileCacheForNonLinux = new(sync.Map)
		// 文件创建缓存处理
		go func() {
			for {
				receiver.newFileCacheForNonLinux.Range(func(key, value interface{}) bool {
					filePath := key.(string)
					fileSize := value.(int64)
					size := util.FileUtil.FileSize(filePath)
					if size == fileSize {
						receiver.newFileCacheForNonLinux.Delete(key)
						receiver.fileHandlerChannel <- &fileEvent{
							FilePath: filePath,
							Op:       Created,
						}
						if receiver.EnableDebugLog {
							logger.Sugar.Debugf("新文件：%s 大小：%d\n", filePath, size)
						}
					} else {
						receiver.newFileCacheForNonLinux.Store(filePath, size)
					}
					return true
				})
				time.Sleep(time.Duration(time.Millisecond.Nanoseconds() * receiver.WriteTime))
			}
		}()

	}
	var err error
	receiver.watcher, err = fsnotify.NewWatcher()
	if err != nil {
		logger.Sugar.Error(err)
	}
	defer func() {
		_ = receiver.watcher.Close()
	}()

	go func() {
		for {
			select {
			case event, ok := <-receiver.watcher.Events:
				if !ok {
					logger.Sugar.Error("获取目录监听事件通道失败")
					time.Sleep(time.Second)
					continue
				}

				// 监听创建事件
				if event.Op&fsnotify.Create == fsnotify.Create || event.Op&fsnotify.Rename == fsnotify.Rename {
					if !file.IsFile(event.Name) {
						dir := event.Name
						err := receiver.watcher.Add(dir)
						if err != nil {
							logger.Sugar.Error("新目录：", dir, "，监听失败", err)
						} else {
							logger.Sugar.Info("成功监听目录：", dir)
						}
					} else {
						if !receiver.EnableFileCreateHandler {
							continue
						}
						receiver.newFileCacheForNonLinux.Store(event.Name, util.FileUtil.FileSize(event.Name))
					}
				}

				// 监听写操作
				if event.Op&fsnotify.Write == fsnotify.Write {
					if !receiver.EnableFileWriteHandler {
						continue
					}
					if receiver.EnableDebugLog {
						logger.Sugar.Debugf("文件写入：%s\n", event.Name)
					}
					receiver.fileHandlerChannel <- &fileEvent{
						FilePath: event.Name,
						Op:       Write,
					}
				}

				// 监听删除操作
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					if !receiver.EnableFileDelHandler {
						continue
					}
					if receiver.EnableDebugLog {
						logger.Sugar.Debugf("删除文件：%s\n", event.Name)
					}
					receiver.fileHandlerChannel <- &fileEvent{
						FilePath: event.Name,
						Op:       Remove,
					}
				}
			case err, ok := <-receiver.watcher.Errors:
				logger.Sugar.Error("目录监听错误 ", err, ok)
				if !ok {
					logger.Sugar.Error("获取目录监听事件通道失败")
					time.Sleep(time.Second)
				}
			}
		}
	}()

	var done = make(chan bool)

	var listeningCount int32
	for _, dir := range watchDirs {
		dir := strings.TrimSpace(dir)
		if dir == "" {
			continue
		}
		homeDir, e := os.UserHomeDir()
		if e != nil {
			logger.Sugar.Error(e)
		}
		if strings.Contains(dir, "~") {
			dir = strings.Replace(dir, "~", homeDir, 1)
		}
		if file.IsExist(dir) {
			if file.IsFile(dir) {
				logger.Sugar.Error("路径：", dir, "，是一个文件，无法监听")
			} else {
				receiver.fsnotifyListening(dir, &listeningCount)
			}
		} else {
			logger.Sugar.Error("目录：", dir, "，不存在，监听失败")
		}

	}
	if listeningCount == 0 {
		logger.Sugar.Warn("未监听任何目录")
		done <- true
	} else {
		logger.Sugar.Info("共监听目录 ", listeningCount, " 个")
	}

	<-done
}
