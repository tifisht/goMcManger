package utils

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
)

var (
	// 定义命令参数
	cmd = exec.Command("mcdrefored.exe")
	//cmd = exec.Command("ping", "-t", "localhost")

	//标准输入输出接口缓冲区
	Stdout, Stderr       bytes.Buffer
	ErrStdout, ErrStderr error

	stdoutIn, _ = cmd.StdoutPipe()
	stderrIn, _ = cmd.StderrPipe()

	stdin, _ = cmd.StdinPipe()
	//开启状态
	IfStart = false
)

type Command struct {
	Context string `json:"context"`
}

// 启动函数
func StartServer() error {

	if err := cmd.Start(); err != nil {
		return err
	}
	//改服务器状态
	IfStart = true

	//将输出传递到缓冲
	go func() {
		_, ErrStdout = io.Copy(&Stdout, stdoutIn)
		io.Copy(os.Stdout, &Stdout)
	}()

	go func() {
		_, ErrStderr = io.Copy(&Stderr, stderrIn)
	}()

	//等待子进程结束
	err := cmd.Wait()

	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	IfStart = false

	return nil
}

//子进程输入函数

func Input(cmd string) {
	io.WriteString(stdin, cmd)
}
