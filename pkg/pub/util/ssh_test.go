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
 * long.qian 2021-10-20 09:43 创建
 */

/**
 * @author long.qian
 */

package util

import (
	"fmt"
	"testing"
)

func Test_sshUtil_GetAllPidByProcessSign(t *testing.T) {
	fmt.Println(SshUtil.GetAllPidByProcessSign(SshConfig{
		LoginName: "qianlong",
		LoginPswd: "qianlong",
		Host:      "192.168.0.121",
		Port:      22,
	}, "ffmpeg"))
}

