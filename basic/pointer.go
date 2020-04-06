package main

import "fmt"

func main() {
	//define a array with values
	var a = [3]int{1, 2, 3}

	for i, x := range a {
		fmt.Printf("No.%d is %d\n", i, x)
	}

	//define a array without values
	var b [3]int

	//look at the way "len()" to get array's length
	fmt.Println("length of b is ", len(b))

	var c = [3]int{100, 200, 300}
	var ptr *int

	for i := 0; i < 3; i++ {
		// & 符号会生成一个指向其作用对象的指针。
		ptr = &c[i]
		// * 符号表示指针指向的底层的值。
		// 与 C 不同，Go 没有指针运算。
		fmt.Printf("the address of c[%d] is %d, its value is %d\n", i, ptr, *ptr)
	}

}
