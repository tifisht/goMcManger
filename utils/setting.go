package utils

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Admin  string
	Secret string
)

type Login struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func init() {
	//读取文件内容
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件有误，请检查文件路径")
		os.Exit(1)
	}

	LoadServer(file)
	LoadData(file)
}

// 服务器启动
func LoadServer(file *ini.File) {

	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
}

// 数据初始化
func LoadData(file *ini.File) {
	Admin = file.Section("administration").Key("admin").MustString("root")
	Secret = file.Section("administration").Key("root").MustString("ttserverWeb")
}
