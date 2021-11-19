package util

import (
	"bytes"
	"fmt"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/logger"
	"golang.org/x/crypto/ssh"
	"net"
	"strings"
	"time"
)

var (
	SshUtil = new(sshUtil)
)

type sshUtil struct {

}

type SshConfig struct {
	LoginName string
	LoginPswd string
	Host      string
	Port      int
}

type SshResult struct {
	IsConnect bool   `json:"isConnect"`
	Result    string `json:"result"`
	Err       error  `json:"err"`
}

// IsSafetyCmd 判断命令是否是允许的安全命令
func (me *sshUtil) IsSafetyCmd(cmd string) (string,bool) {
	if strings.Contains(cmd, "rm") {
		if len(strings.Split(cmd, "/")) <= 1 {
			return fmt.Sprintf("rm命令 %s 不能删除小于2级的文件",cmd),false
		}
	}
	return "",true
}

// KillForceByProcessSign 根据进程的命令标识，强制 Kill 对应的进程
func (me *sshUtil) KillForceByProcessSign(config SshConfig, sign string) error {
	cmd := "kill -9 `ps aux | grep \"" + sign + "\"| grep -v grep|awk '{print $2}'`"
	r := me.ExecRemoteSshCommand(config, []string{
		cmd,
	})
	return r.Err
}

// GetAllPidByProcessSign 根据进程的命令标识，获取对应的进程 ID，返回对应的进程 ID 集合
func (me *sshUtil) GetAllPidByProcessSign(config SshConfig, sign string) ([]string, error) {
	pidCmd := "ps aux | grep \"" + sign + "\"| grep -v grep|awk '{print $2}'"
	r := me.ExecRemoteSshCommand(config, []string{
		pidCmd,
	})
	if r.Err != nil {
		return nil, r.Err
	}
	var allPid []string
	ss := strings.Split(strings.TrimSpace(r.Result), "\n")
	for _, s := range ss {
		pid := strings.TrimSpace(s)
		if pid != "" {
			allPid = append(allPid, pid)
		}
	}
	return allPid, nil
}

// ExecRemoteSshCommand 执行指定机器的 SSH 命令
func (me *sshUtil) ExecRemoteSshCommand(config SshConfig, cmds []string) SshResult {
	return me.ExecRemoteSshCommandAutoSessionCloseTimeout(config, cmds, 0)
}

// ExecRemoteSshCommandAutoSessionCloseTimeout 执行远程命令，超时时，自动关闭session，sessionCloseSecond=0 代表不设置超时强制关闭session
func (me *sshUtil) ExecRemoteSshCommandAutoSessionCloseTimeout(config SshConfig, cmds []string, sessionCloseSecond int64) SshResult {
	var session *ssh.Session
	client, err := me.GetSshClient(config.LoginName, config.LoginPswd, config.Host, config.Port)
	if err != nil {
		return SshResult{
			IsConnect: false,
			Err:       err,
		}
	}
	// create session
	if session, err = client.NewSession(); err != nil {
		return SshResult{
			IsConnect: false,
			Err:       fmt.Errorf("Failed New SSH Session To %s : %s", config.Host, err.Error()),
		}
	}

	defer me.closeClient(session,client)

	go func() {
		if sessionCloseSecond > 0 {
			sleep, err := time.ParseDuration(fmt.Sprintf("%ds", sessionCloseSecond))
			if err != nil {
				logger.Sugar.Errorf("SSH自动session关闭失败：%s", err.Error())
				return
			}
			time.Sleep(sleep)
			me.closeClient(session,client)
		}
	}()

	command := ""

	for i, cmd := range cmds {
		if msg,ok := me.IsSafetyCmd(cmd);!ok {
			return SshResult{
				IsConnect:true,
				Err:fmt.Errorf(msg),
			}
		}
		if i == 0 {
			command = cmd
		} else {
			command = command + " && " + cmd
		}
	}

	var e bytes.Buffer
	var b bytes.Buffer
	session.Stdout = &b
	session.Stderr = &e
	if command != "" {
		if err := session.Run(command); err != nil {
			return SshResult{
				IsConnect: true,
				Err:       fmt.Errorf("命令 %s 执行失败：%s\n%s", command, err.Error(), e.String()),
				Result:    e.String(),
			}
		}
		logger.Sugar.Info(command)
	}
	return SshResult{
		Result:    b.String(),
		Err:       nil,
		IsConnect: true,
	}
}

func (me *sshUtil) closeClient(session *ssh.Session,client *ssh.Client) {
	err := client.Close()
	if err != nil {
		logger.Sugar.Error("SSH Client Close Error",err)
	}
	//必须关闭Client，才能释放该ssh连接句柄
	_ = session.Close()
}

func (me *sshUtil) GetSshClient(user, password, host string, port int) (*ssh.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 8 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connect to ssh
	addr = fmt.Sprintf("%s:%d", host, port)
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, fmt.Errorf("Connect To %s Error : %s", addr, err.Error())
	}
	return client, nil
}
