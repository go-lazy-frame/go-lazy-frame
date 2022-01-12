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
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/util"
	"github.com/toolkits/file"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
	"time"
)

type fileEvent struct {
	FilePath string
	Op       fsnotify.Op
}

// FileWatch 文件监听
type FileWatch struct {
	// 文件记录缓存
	filesRecordCache *sync.Map
	// 文件处理channel
	fileHandlerChannel chan *fileEvent
	// 监听器
	watcher *fsnotify.Watcher
	// 判断新文件是否写完的时间更新阈值，默认 500（根据网络环境，自行调整），单位毫秒
	WriteTime int64
	// 是否允许 Debug 日志输出，默认 false
	EnableDebugLog bool
	// 是否处理文件创建事件， 默认 false
	EnableFileCreateHandler bool
	// 是否处理文件删除事件，默认 false
	EnableFileDelHandler bool
}

// StartFileWatch 开始目录监听
func (receiver *FileWatch) StartFileWatch(fileHandler func(filePath string, op fsnotify.Op), watchDirs ...string) {
	if receiver.WriteTime == 0 {
		receiver.WriteTime = 500
	}
	receiver.filesRecordCache = new(sync.Map)
	receiver.fileHandlerChannel = make(chan *fileEvent, runtime.NumCPU()*2)
	go func() {
		go func() {
			for {
				e := <-receiver.fileHandlerChannel
				go fileHandler(e.FilePath, e.Op)
			}
		}()
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
						if !receiver.EnableFileCreateHandler {
							continue
						}
						if !file.IsFile(event.Name) {
							dir := event.Name
							err := receiver.watcher.Add(dir)
							if err != nil {
								logger.Sugar.Error("新目录：", dir, "，监听失败", err)
							} else {
								logger.Sugar.Info("成功监听目录：", dir)
							}
						} else {
							receiver.filesRecordCache.Store(event.Name, util.TimeUtil.GetMilliTime(time.Now()))
							if receiver.EnableDebugLog {
								logger.Sugar.Debugf("发现新文件：%s，正在等待其写入完成...\n", event.Name)
							}
							go func() {
								for {
									if t, ok := receiver.filesRecordCache.Load(event.Name); ok {
										// 若一段时间后，文件没有任何写操作，则认为该文件已传输完毕
										sub := util.TimeUtil.GetMilliTime(time.Now()) - t.(int64)
										if sub >= receiver.WriteTime {
											receiver.filesRecordCache.Delete(event.Name)
											receiver.fileHandlerChannel <- &fileEvent{
												FilePath: event.Name,
												Op:       fsnotify.Create,
											}
											break
										}
									}
								}
							}()
						}
					}

					// 监听写操作
					if event.Op&fsnotify.Write == fsnotify.Write {
						if !receiver.EnableFileCreateHandler {
							continue
						}
						receiver.filesRecordCache.Store(event.Name, util.TimeUtil.GetMilliTime(time.Now()))
					}

					// 监听删除操作
					if event.Op&fsnotify.Remove == fsnotify.Remove {
						if !receiver.EnableFileDelHandler {
							continue
						}
						receiver.fileHandlerChannel <- &fileEvent{
							FilePath: event.Name,
							Op:       fsnotify.Remove,
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
					receiver.listening(dir, &listeningCount)
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
	}()
}

func (receiver *FileWatch) listening(adsPath string, listeningCount *int32) {
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
					receiver.listening(path.Join(adsPath, d.Name()), listeningCount)
				}
			}
		}
	}
}
