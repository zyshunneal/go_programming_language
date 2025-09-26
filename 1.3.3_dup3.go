package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t %s \n", n, line)
		}
	}

}

/*
➜  go_programming_language git:(main) ✗ go run 1.3.3_dup3.go tt.txt
3        zyshun
2        neal
2
dup 的前两个版本以"流”模式读取输入，并根据需要拆分成多个行。理论上，这些程序可以处理任意数量的输入数据。
还有另一个方法，就是一口气把全部输入数据读到内存中，一次分割为多行，然后处理它们。下面这个版本，dup3，
就是这么操作的。这个例子引入了 ReadFile 函数（来自于io/ioutil包），其读取指定文件的全部内容，strings.Split
函数把字符串分割成子串的切片。（Split 的作用与前文提到的 strings.Join 相反。请注意如果在Windows下测试注意换行是否为\r\n，
否则最后一行是否有空行将会影响结果。）

我们略微简化了 dup3。首先，由于 ReadFile 函数需要文件名作为参数，因此只读指定文件，不读标准输入。其次，
由于行计数代码只在一处用到，故将其移回 main 函数。
*/
