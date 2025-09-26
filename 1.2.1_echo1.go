package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

/*
大多数的程序都是处理输入和输出； 这也正是计算 的 定义。
通常情况下输入来自于程序外部：文件、网络链接、其他程序的输出
敲键盘的用户、命令行参数或其它类似的输入源。
本节：命令行参数
os.Args() 变量是一个字符串的切片，切片是 Go 语言的基础概念。
os.Args() 的第一个元素 os.Args[0] 是命令本身的名字。
其它的元素则是程序启动时传给它的参数。
本节是 echo 的一份实现，echo 把 它的命令行参数打印成一行。
*/
