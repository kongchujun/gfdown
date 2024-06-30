// Package server /*
/*
@File: ip
@Time: 2024/6/30-12:58
@Author: Lewis Kong
@Email: kongplusjun@163.com
@Desc: Some description of the file
*/
package server

import (
	"fmt"
	"net"
)

// 获取本地IP地址
func getLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String(), nil
			}
		}
	}

	return "", fmt.Errorf("没有找到局域网IP地址")
}
