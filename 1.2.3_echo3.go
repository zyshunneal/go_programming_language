package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, arg := range os.Args[:] {
		fmt.Printf("%d %s \n", idx, arg)
	}

}

/*
练习 1.1： 修改 echo 程序，使其能够打印 os.Args[0]，即被执行命令本身的名字。

➜  go_programming_language git:(main) ✗ go run  1.1.3_echo3.go a b c
0 /var/folders/x4/n2gq4tq928bf34slb0brgy3m0000gn/T/go-build3536104312/b001/exe/1.1.3_echo3
1 a
2 b
3 c
*/
