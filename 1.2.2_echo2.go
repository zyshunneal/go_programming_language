package main

import (
	"fmt"
	"os"
)

func main() {
	s, seq := "", ""
	for _, arg := range os.Args[1:] {
		s += seq + arg
		seq = " "
	}

	fmt.Println(s)

}

/*
for 循环的另一种形式， 在某种数据类型的区间(range) 上遍历，如字符串、切片。
这里用 rang 实现 echo 的另一个版本

每次循环迭代，range 产生一对值；索引以及在该索引处的元素值。这个例子不需要索引，
但 range 的语法要求，要处理元素，必须处理索引。一种思路是把索引赋值给一个临时变
量（如 temp）然后忽略它的值，但 Go 语言不允许使用无用的局部变量（local variables），
因为这会导致编译错误。

Go 语言中这种情况的解决方法是用 空标识符（blank identifier），即 _（也就是下划线）。
空标识符可用于在任何语法需要变量名但程序逻辑不需要的时候（如：在循环里）丢弃不需要的
循环索引，并保留元素值。大多数的 Go 程序员都会像上面这样使用 range 和 _ 写 echo
程序，因为隐式地而非显式地索引 os.Args，容易写对。
➜  go_programming_language git:(main) ✗ go run 1.1.2_echo2.go a b c d
a b c d
*/
