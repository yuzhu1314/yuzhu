package main

import (
	"fmt"
	"math"
)

// 6.编写一个函数，求100以内的质数
// 质数又称素数。一个大于1的自然数，除了1和它自身外，不能被其他自然数整除的数叫做质数；否则称为合数
func isPrime(i int) bool {
	for j := 2; float64(j) <= math.Sqrt(float64(i)); j++ {
		if i%j == 0 {
			return false

		}
	}
	return true
}
func main() {
	var str []int //存放质数
	j := 0
	for i := 2; i < 100; i++ {
		if isPrime(i) {
			str = append(str, i)
			j++
		}
	}
	fmt.Println(str)
}

//结果
//[2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97]
