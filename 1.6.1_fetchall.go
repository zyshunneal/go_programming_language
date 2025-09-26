package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f %7d %s", secs, nbytes, url)
}

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

/*
➜  go_programming_language git:(main) ✗ go build 1.6.1_fetchall.go
➜  go_programming_language git:(main) ✗ ./1.6.1_fetchall http://www.baidu.com http://www.sougou.com
0.07    2381 http://www.baidu.com
0.67   19515 http://www.sougou.com
0.67s elapsed
Go 语言最有意思并且最新奇的就是对并发编程的支持。
并发编程是一个大话题。后续专门学习
*/
