// Package server /*
/*
@File: fileserver
@Time: 2024/6/30-12:57
@Author: Lewis Kong
@Email: kongplusjun@163.com
@Desc: Some description of the file
*/
package server

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

// StartFileServer 启动文件服务器
func StartFileServer(dir, username, password, port string) error {
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return fmt.Errorf("无法解析目录: %v", err)
	}

	// 获取本地IP地址
	localIP, err := getLocalIP()
	if err != nil {
		return fmt.Errorf("无法获取本地IP地址: %v", err)
	}

	// 创建一个文件服务器
	fs := http.FileServer(http.Dir(absDir))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if username != "" && password != "" {
			if !checkAuth(w, r, username, password) {
				return
			}
		}
		fs.ServeHTTP(w, r)
	})

	log.Printf("文件服务器正在运行，目录：%s\n", absDir)
	log.Printf("访问地址：http://%s:%s\n", localIP, port)
	return http.ListenAndServe(":"+port, nil)
}
