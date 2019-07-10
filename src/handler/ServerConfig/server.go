package main

import (
	"net/http"
)

//可通过Server结构对服务器进行更详细的配置
func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil}
	server.ListenAndServe()
}

//Server结构
// type Server struct {
// 	Addr			string
// 	Handler			Handler
// 	ReadTiemout		time.Duration
// 	WriteTimeout	time.Duration
// 	MaxHeaderBytes	int
// 	TLSConfig		*tls.Config
// 	TLSNextProto	map[string]func(*Server, *tls.Conn, Handler)
// 	ConnState		func(net.Conn, ConnState)
// 	ErrorLog		*log.Logger
// }
