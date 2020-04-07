package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// `if` 语句可以在条件之前执行一个简单的语句。
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func switchfunc() {
	var x interface{}
	// i.(type) 只能在switch中使用，函数没有返回值
	switch i := x.(type) {
	case nil:
		fmt.Println(" x 的类型 :%T", i)
	case int:
		fmt.Println("x 是 int 型")
	case float64:
		fmt.Println("x 是 float64 型")
	case func(int) float64:
		fmt.Println("x 是 func(int) 型")
	case bool, string:
		fmt.Println("x 是 bool 或 string 型")
	default:
		fmt.Println("未知型")
	}

	// 没有条件的 switch 同 `switch true` 一样。
	// 这一构造使得可以用更清晰的形式来编写长的 if-then-else 链。
	//注意fallthrough的用法
	//switch语句一般会默i认最后一句话是带有break，
	//但使用fallthrough之后，即使下一个case是false，也会强制执行下一个case
	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
}

func main() {

	// 注意第二个参数后面还有一个逗号
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	switchfunc()
}
