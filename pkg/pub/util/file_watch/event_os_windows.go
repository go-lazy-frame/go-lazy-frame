//go:build windows

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
 * long.qian 2022-01-17 15:20 创建
 */

/**
 * @author long.qian
 */

package file_watch

func (receiver *FileWatch) eventHandler(watchDirs []string) {
	receiver.nonLinuxEventHandler(watchDirs)
}
