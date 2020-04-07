package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
	name string
}

func (nokiaPhone NokiaPhone) call() {
	nokiaPhone.name = "nokia call"
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
	name string
}

// 区别于上一个方法，这个地方传递的是指针
func (iPhone *IPhone) call() {
	iPhone.name = "iphone call"
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var phone Phone

	nokia := new(NokiaPhone)
	phone = nokia // 注意这里是有转型的
	phone.call()
	println(nokia)

	iphone := new(IPhone)
	// 按照http://go-tour-zh.appspot.com/methods/4
	// 这个地方应该是 phone = &iphone
	// 为什么不报错？
	phone = iphone
	phone.call()
	fmt.Println(iphone)

}
