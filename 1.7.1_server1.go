package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
Go 语言的内置库使得写一个类似的 fetch 的 web 服务器变得异常简单。
在 server1 中我们会实现服务器返回当前用户正在访问的 url。
*/

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q \n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
