package main

import (
	"fmt"
)

// Go 闭包
// 闭包是一个函数值，它来自函数体的外部的变量引用。
// 函数可以对这个引用值进行访问和赋值；换句话说这个函数被“绑定”在这个变量上。
// 匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
// 例如，函数 adder 返回一个闭包。每个闭包都被绑定到其各自的 sum 变量上
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	now, next := 0, 1
	return func() int {
		x := now
		now, next = next, now+next
		return x
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 5; i++ {
		fmt.Println(
			pos(i),
			neg(-i),
		)
	}

	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(fib(), " ")
	}
}
