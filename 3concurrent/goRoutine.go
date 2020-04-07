package main

import (
	"fmt"
	"runtime"
	"time"
)

func goRoutine1() {
	for i := 0; i < 5; i++ {
		fmt.Println("goRoutine1 is runing ", i)
		time.Sleep(time.Second * 1)
	}
}

func goRoutine2() {
	for i := 0; i < 5; i++ {
		fmt.Println("goRoutine2 is runing ", i)
		time.Sleep(time.Second * 1)
	}
}

func differance1() {
	fmt.Println("In differance1()")
	// set the number of cpu cores
	runtime.GOMAXPROCS(1)

	// to get the number of the CPU cores
	// var num = runtime.NumCPU

	// to get the number of goroutines
	// var goRoutineNum = runtime.NumGoroutine()
	for i := 0; i < 10; i++ {
		// The compiler packages the parameters and functions that follow go
		// into objects and waits for system scheduling.
		// so, there the object is { println, current_i }
		go fmt.Printf("%d ", i) // the result is: 9 0 1 2 3 4 5 6 7 8
	}

	// if the 'goroutine' find the current condition is not adequote,
	// release the occupation of cpu by using release function
	runtime.Gosched()

	time.Sleep(time.Second * 1)
	fmt.Println()
}

func differance2() {
	fmt.Println("In differance2()")
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		// The compiler packages the parameters and functions that follow go
		// into objects and waits for system scheduling.
		// so, there the object is { main.func_xxx, nil }
		go func() {
			fmt.Printf("%d ", i) // the result is: 10 10 10 10 10 10 10 10 10 10
		}()
	}

	runtime.Gosched()

	time.Sleep(time.Second * 1)
	fmt.Println()
}

// to handle the goroutine's error
func addele(a []int, i int) {
	// defer后面的函数在defer语句所在的函数执行结束的时候会被调用；
	defer func() { // 匿名函数捕捉错误
		err := recover()
		if err != nil {
			fmt.Println("add ele fail")
		}
	}()
	a[i] = i
	fmt.Println(a)
}

func main() {
	fmt.Println("In main()")

	// you will see that the call of goR1 and goR2 is randomly
	// because the goroutines run asynchronously
	// 首先了解：进程、线程、协程之间地联系与区别，goroutine的核心是协程地并行计算
	go goRoutine1()
	goRoutine2()

	// know something about keyword 'go'
	// differance1()
	// differance2()

	// to wait the back of goRoutine1 and goRoutine2
	// because goroutine runs asynchronously, main() may end earlier than goroutines
	// if main returns, the call of goR1 and goR2 will return too
	time.Sleep(time.Second * 6)
}
