package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Manning publications Co",
		HttpOnly: true,
	}
	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())

	//直接使用SetCookie方法
	// http.SetCookie(w, &c1)
	// http.SetCookie(w, &c2)
}

//获取cookie
func getCookie(w http.ResponseWriter, r *http.Request) {
	//r.Header["Cookie"]返回一个切片，包含了多个cookie
	// h := r.Header["Cookie"]
	// fmt.Fprintln(w, h)

	//获取指定Cookie
	cl, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "error")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, cl)
	fmt.Fprintln(w, cs)

}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}
