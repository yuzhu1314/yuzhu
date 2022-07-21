package main

import (
	"fmt"
)

//3.回文数的判断
//回文数的概念：即是给定一个数，这个数顺读和逆读都是一样的。例如：121，1221是回文数，123，1231不是回文数。

func isHuiWen(a string) bool {
	j := len(a) - 1
	b := true
	for i := 0; i <= len(a)/2-1; i++ {
		if a[i] != a[j] {
			b = false
		}
		j--
	}
	return b
}
func main() {
	a := "121"
	result := isHuiWen(a)
	fmt.Println(result)
}
