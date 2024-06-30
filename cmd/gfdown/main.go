// Package gfdown /*
/*
@File: main.go
@Time: 2024/6/30-12:57
@Author: Lewis Kong
@Email: kongplusjun@163.com
@Desc: Some description of the file
*/
package main

import (
	"flag"
	"gfdown/internal/server"
	"log"
)

func main() {
	// 命令行参数解析
	var dir, username, password, port string
	flag.StringVar(&dir, "dir", ".", "要暴露的目录")
	flag.StringVar(&username, "username", "", "登录用户名")
	flag.StringVar(&password, "password", "", "登录密码")
	flag.StringVar(&port, "port", "8666", "服务器端口")
	flag.Parse()

	// 启动文件服务器
	if err := server.StartFileServer(dir, username, password, port); err != nil {
		log.Fatalf("无法启动文件服务器: %v", err)
	}
}
