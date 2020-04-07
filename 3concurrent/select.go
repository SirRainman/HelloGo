package main

import (
	"fmt"
	"time"
)

/*
	select 语句使得一个 goroutine 在多个通讯操作上等待。
	select 会阻塞，直到条件分支中的某个可以继续执行，这时就会执行那个条件分支。
	当多个都准备好的时候，会随机选择一个。
	当 select 中的其他条件分支都没有准备好的时候，`default` 分支会被执行。
*/

func fib(ch chan int, isFinished chan bool) {
	now, next := 0, 1
	for {
		select {
		case ch <- now:
			now, next = next, now+next
		case <-isFinished:
			fmt.Println("finished")
			return
		default:
			fmt.Println("wait for read")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ch := make(chan int)
	isFinished := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		isFinished <- true
	}()

	fib(ch, isFinished)
}
