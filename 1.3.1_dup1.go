package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t %s\n", n, line)
		}
	}
}

/*
对文件的拷贝、打印、搜索、排序、统计或类似的事情的程序都有一个差不多的结构：一个处理输入的循环，
在每个元素上执行计算处理，在处理的同时或最后产出输出。
dup 的第一个版本打印标准输入中多次出现的行，以重读次数开头。

➜  go_programming_language git:(main) ✗ cat tt.txt
zyshun
neal
zyshun
neal
zyshun
zz

➜  go_programming_language git:(main) ✗ cat tt.txt | go run 1.3.1_dup1.go
3        zyshun
2        neal
*/
