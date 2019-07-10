package main

import (
	"fmt"
	"net/http"
)

//如果没有为服务器编写任何处理器，则会返回404，如simplest里的server文件
//一个处理器就是一个拥有ServeHTTP方法的接口，这个方法需要ResponseWrite接口和Request指针
//ListenAndServe默认是DefaultServeMux，它拥有ServeHTTP方法，不仅是多路复用器，还是处理器
//DefaultServeMux是个特殊的处理器，唯一要做的就是根据请求的URL定向到不同的处理器

//MyHandler 自定义的handler处理器
type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	server.ListenAndServe()
}
