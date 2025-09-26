package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)

	}
}

/*
对于很多现代应用来说，访问互联网上的信息和访问本地文件系统一样重要。Go 语言在 net
这个强大 package 的帮助下提供了一系列的 package 来做这个事情，使用这个包可以更简单
的用网络收发信息，还可以建立更底层的网络链接，编写服务器程序。

*/
