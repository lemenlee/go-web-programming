package main

import (
	"fmt"
	"net/http"
)

//如果一个键同时有表单键值对和URL键值对，可以访问Request结构的PostForm字段来获得表单键值对
func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//r.Form同时获取键值对和URL键值对
	// fmt.Fprintln(w, r.Form)

	// println(r.PostForm)

	//解析 multipart/form-data格式数据
	// r.ParseMultipartForm(1024)
	// fmt.Fprintln(w, r.MultipartForm)

	//FormValue方法直接访问给定键的值，且会自动调用PareseForm
	fmt.Fprintln(w, r.FormValue("hello"))
	fmt.Fprintln(w, r.Form)

}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
