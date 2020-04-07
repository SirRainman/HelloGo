package main

import "fmt"

func max(num1, num2 int) int {
	var result int
	if num1 < num2 {
		result = num2
	} else {
		result = num1
	}
	return result
}

// go can return serveral values
// 当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。
func swap(str1, str2 string) (string, string) {
	return str2, str1
}

func split(sum int) (x, y int) {
	// 有意思，为什么会这样呢？
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {

	// Go 的基本类型
	// bool
	// string

	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte // uint8 的别名
	// rune // int32 的别名
	// 	// 代表一个Unicode码

	// float32 float64
	// complex64 complex128

	var a int = 1
	var b int = 2

	var ret = max(a, b)
	fmt.Printf("the bigger is %d\n", ret)

	// 注意这个format
	const f = "%T(%v)\n"
	fmt.Printf(f, "type", "value")

	//be care that multi returned values
	str1, str2 := swap("good", "bad")
	fmt.Println("the rusult swap(str1, str2): ", str1, str2)

	fmt.Printf("the reuslt split(18): ")
	fmt.Println(split(18))

	// 如果初始化是使用表达式，则可以省略类型；变量从初始值中获得类型。
	// var c, python, java = true, 2, "no!"
	// 在函数中，`:=` 简洁赋值语句在明确类型的地方，可以用于替代 var 定义。
	// 函数外的每个语句都必须以关键字开始（`var`、`func`、等等），`:=` 结构不能使用在函数外。
	c, python, java := true, 2, "no!"
	fmt.Println(c, python, java)

}
