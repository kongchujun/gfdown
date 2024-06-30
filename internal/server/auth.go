// Package server /*
/*
@File: auth
@Time: 2024/6/30-12:57
@Author: Lewis Kong
@Email: kongplusjun@163.com
@Desc: Some description of the file
*/
package server

import (
	"net/http"
)

// 检查基本认证
func checkAuth(w http.ResponseWriter, r *http.Request, username, password string) bool {
	user, pass, ok := r.BasicAuth()
	if !ok || user != username || pass != password {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized.\n"))
		return false
	}
	return true
}
