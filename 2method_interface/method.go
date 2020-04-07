package main

import "fmt"

type Square struct {
	long, width float64
}

// Go 没有类。然而，仍然可以在结构体类型上定义方法。
// 可以对包中的 任意 类型定义任意方法，而不仅仅是针对结构体。
// 但是，不能对来自其他包的类型或基础类型定义方法。

// 方法接收者 出现在 func 关键字和方法名之间的参数中。
func (s *Square) Area() float64 {
	// s.long = -1 如果这里对s进行修改的的话，因为是传递的指针，因此对原有的数据是有改变的
	return s.long * s.width
}

type MyFloat float64

// 问：为什么这里不是指针，但是上面是指针？http://go-tour-zh.appspot.com/methods/3
// 有两个原因需要使用指针接收者。
// 		首先避免在每个方法调用该结构体副本的值（如果值类型是大的结构体的话会更有效率），省略掉了拷贝那一步。
// 		其次，方法可以修改接收者指向的值，即可以修改指针指向的结构体里面的数据。如果是值传递的话，修改的是拷贝后结构体的数据。
func (f MyFloat) Abs() float64 {
	f = 1 //这里更新的只是副本的值，不会对原有数据进行更改
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	s := &Square{2, 3}
	fmt.Println(s.Area())
}
