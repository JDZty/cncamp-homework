// 接收客户端 request，并将 request 中带的 header 写入 response header
// 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
// Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
// 当访问 localhost/healthz 时，应返回 200
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	initHandler()

	fmt.Println("================== server already start ==================")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("server start error:", err)
	}

}

func initHandler() {
	http.HandleFunc("/reqh_to_resph", reqHToRespH)
	http.HandleFunc("/sys_version", sysVersion)
	http.HandleFunc("/log_to_console", logToConsole)
	http.HandleFunc("/healthz", healthz)
}

// 接收客户端 request，并将 request 中带的 header 写入 response header
func reqHToRespH(resp http.ResponseWriter, req *http.Request) {

	reqH := req.Header
	fmt.Printf("request hearder is %v\n", reqH)
	for k, v := range reqH {
		resp.Header().Add(k, strings.Join(v, ","))
	}

}

// 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
func sysVersion(resp http.ResponseWriter, req *http.Request) {
	v := os.Getenv("VERSION")
	resp.Header().Add("version", v)
}

// Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
func logToConsole(resp http.ResponseWriter, req *http.Request) {

	clientIp := req.RemoteAddr
	fmt.Println("clinet ip is", clientIp)

	code := 200
	resp.WriteHeader(code)
	fmt.Printf("status code is %d\n", code)

}

// 当访问 localhost/healthz 时，应返回 200
func healthz(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(200)
}
