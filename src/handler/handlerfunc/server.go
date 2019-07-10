package main

import (
	"fmt"
	"net/http"
)

//直接使用handlerfunc，把一个带有正确签名的函数f转换为带有f方法的Handler
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	//还可以这样写
	// hellohandler := http.HandlerFunc(hello)
	// worldhandler := http.HandlerFunc(world)

	// http.Handle("/hello", &hellohandler)
	// http.Handle("/world", &worldhandler)

	server.ListenAndServe()
}
