package main

import (
	"net/http"
)

//cert.pem是SSL证书
//key.pem是服务器私钥
//生产环境通过VeriSign， Thawte， Comodo SSl申请

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil}
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
