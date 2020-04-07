package main

import (
	"fmt"
	"time"
)

func define() {
	// define chan
	// 和 map 与 slice 一样，channel 使用前必须创建：
	readOnlyChan := make(<-chan int, 10)
	writeOnlyChan := make(chan<- int, 10)
	normalChan := make(chan int, 10)

	// note:
	// 管道如果未关闭，在读取超时会则会：引发deadlock异常
	// 管道如果关闭进行写入数据会：pannic
	// 当管道中没有数据时候读取会：读取到默认值，如int类型默认值是0
	// 使用range循环管道，如果管道未关闭会：引发deadlock错误。
	// 如果采用for死循环已经关闭的管道，当管道没有数据时候，读取的数据会：是管道的默认值，并且循环不会退出。
	// 默认情况下，在另一端准备好之前，发送和接收都会阻塞。
	// 这使得 goroutine 可以在没有明确的锁或竞态变量的情况下进行同步。

	// operation:
	var x = <-readOnlyChan // read data
	writeOnlyChan <- x     // write data
	y, ok := <-normalChan  // read data with state

	fmt.Println(x, y, ok)
}

func operation() {
	ch := make(chan int, 10)

	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)

	fmt.Println("after 'insert', len(channel) is ", len(ch))

	for v := range ch {
		fmt.Println("read from ch: ", v)
	}
	fmt.Println("after 'range', len(channel) is ", len(ch))
}

func chanBuffer_insert(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println(i, " is inserted into ch")
	}
}

func chanBuffer() {
	// channel 可以是 _带缓冲的_。为 make 提供第二个参数作为缓冲长度来初始化一个缓冲
	// ch := make(chan int) // 非缓冲, 只能存一个数据，并且只有当该数据被取出时候才能存下一个数据。
	ch := make(chan int, 6) // 缓冲, 类似一个队列，只有队列满了才可能发送阻塞

	go chanBuffer_insert(ch)

	for j := 0; j < 10; j++ {
		time.Sleep(time.Second * 3)
		fmt.Println(<-ch, "got from ch")
	}
}

func doTask(taskChan chan int, resultChan chan int, exitChan chan bool) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("doTask error:", err)
			return
		}
	}()

	for t := range taskChan {
		fmt.Println("do task", t)
		resultChan <- t
	}

	exitChan <- true
}

func taskPool() {
	taskChan := make(chan int, 20)
	resultChan := make(chan int, 20)
	exitChan := make(chan bool, 5)

	// add tasks
	go func() {
		for i := 0; i < 10; i++ {
			taskChan <- i
		}
		close(taskChan) //发送者可以 close 一个 channel 来表示再没有值会被发送了。
	}()
	/*
		注意： 只有发送者才能关闭 channel，而不是接收者。
		向一个已经关闭的 channel 发送数据会引起 panic。
		注意： channel 与文件不同；
		通常情况下无需关闭它们。只有在需要告诉接收者没有更多的数据的时候才有必要进行关闭，例如中断一个 `range`。
	*/

	// do tasks
	for i := 0; i < 5; i++ {
		go doTask(taskChan, resultChan, exitChan)
	}

	go func() {
		// what does exitChan do??
		// answer: to wait for doTask() finished
		for i := 0; i < 5; i++ {
			<-exitChan
		}
		close(resultChan)
		close(exitChan)
	}()

	for res := range resultChan {
		fmt.Println(res, "is Done")
	}
}

func main() {
	// define()

	// operation()

	chanBuffer()

	// taskPool()
}
