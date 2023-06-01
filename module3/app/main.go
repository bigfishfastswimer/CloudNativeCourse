package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", generalHandler)
	http.HandleFunc("/healthz", healthzHandler)
	//Listen and serve on localhost:80
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
	log.Println("Server started on localhost:80")
}

// create a http server
func generalHandler(w http.ResponseWriter, r *http.Request) {
	//接收客户端 request，并将 request 中带的 header 写入 response header
	for k, v := range r.Header {
		for _, value := range v {
			w.Header().Add(k, value)
		}
	}
	version := os.Getenv("VERSION")
	//读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	w.Header().Set("Version", version)
	w.WriteHeader(http.StatusOK)
	// Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	log.Printf("client ip: %s. http status: %d", r.RemoteAddr, http.StatusOK)
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	// 当访问 localhost/healthz 时，应返回 200
	w.WriteHeader(http.StatusOK)
	log.Printf("api request on healthz/ on client ip: %s. http status: %d", r.RemoteAddr, http.StatusOK)
	// write response header to the client
	io.WriteString(w, "200 OK")


}
