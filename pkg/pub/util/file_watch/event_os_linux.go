//go:build linux

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
 * long.qian 2022-01-17 15:15 创建
 */

/**
 * @author long.qian
 */

package file_watch

import (
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/logger"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/util"
	"github.com/rjeczalik/notify"
	"github.com/toolkits/file"
	"path"
)

func (receiver *FileWatch) eventHandler(watchDirs []string) {
	logger.Sugar.Info("Linux 文件事件处理")
	// Make the channel buffered to ensure no event is dropped. Notify will drop
	// an event if the receiver is not able to keep up the sending pace.
	c := make(chan notify.EventInfo, 100)

	var watched []string
	for _, dir := range watchDirs {
		dir = path.Join(dir, "...")
		if !util.ArrayUtil.IsExistStringArray(&watched, dir) {
			// Set up a watchpoint listening for inotify-specific events within a
			// current working directory. Dispatch each InCloseWrite and InMovedTo
			// events separately to c.
			if err := notify.Watch(dir, c, notify.All, notify.InCloseWrite, notify.InMovedTo); err != nil {
				logger.Sugar.Error("目录：", dir, " 监听失败：", err)
			} else {
				logger.Sugar.Infof("已监听目录：%s\n", dir)
			}
			watched = append(watched, dir)
		}
	}
	defer func() {
		logger.Sugar.Info("停止监听")
		notify.Stop(c)
	}()

	for {
		switch ei := <-c; ei.Event() {
		case notify.Create:
			if !file.IsFile(ei.Path()) {
				dir := ei.Path()
				dir = path.Join(dir, "...")
				if err := notify.Watch(dir, c, notify.All, notify.InCloseWrite, notify.InMovedTo); err != nil {
					logger.Sugar.Error("新目录：", dir, "，监听失败：", err)
				} else {
					logger.Sugar.Info("成功监听目录：", dir)
				}
			}
		case notify.InCloseWrite, notify.InMovedTo:
			if !receiver.EnableFileCreateHandler {
				continue
			}
			filePath := ei.Path()
			receiver.fileHandlerChannel <- &fileEvent{
				FilePath: filePath,
				Op:       Created,
			}
			if receiver.EnableDebugLog {
				logger.Sugar.Debugf("新文件：%s 大小：%d\n", filePath, util.FileUtil.FileSize(filePath))
			}
		case notify.Write:
			if !receiver.EnableFileWriteHandler {
				continue
			}
			if receiver.EnableDebugLog {
				logger.Sugar.Debugf("文件写入：%s\n", ei.Path())
			}
			receiver.fileHandlerChannel <- &fileEvent{
				FilePath: ei.Path(),
				Op:       Write,
			}
		case notify.Remove:
			if !receiver.EnableFileDelHandler {
				continue
			}
			if receiver.EnableDebugLog {
				logger.Sugar.Debugf("删除文件：%s\n", ei.Path())
			}
			receiver.fileHandlerChannel <- &fileEvent{
				FilePath: ei.Path(),
				Op:       Remove,
			}
		}
	}
}
