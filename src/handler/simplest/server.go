package main

import (
	"net/http"
)

//最简化配置，默认是80端口，DefaultServeMux
func main() {
	http.ListenAndServe("", nil)
}
