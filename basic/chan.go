package main 

import (
	"fmt"
	//"time"
)

func define() {
	// define chan
	readOnlyChan := make(<- chan int, 10)
	writeOnlyChan := make(chan <- int, 10)
	normalChan := make(chan int, 10)
	
	// note:
	// 管道如果未关闭，在读取超时会则会引发deadlock异常
	// 管道如果关闭进行写入数据会pannic
	// 当管道中没有数据时候再行读取或读取到默认值，如int类型默认值是0
	// 使用range循环管道，如果管道未关闭会引发deadlock错误。
	// 如果采用for死循环已经关闭的管道，当管道没有数据时候，读取的数据会是管道的默认值，并且循环不会退出。


	// operation:
	var x = <- readOnlyChan // read data
	writeOnlyChan <- x // write data
	y, ok := <- normalChan // read data with state
	
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
		fmt.Println(i, " is inserted into ch")
		ch <- i
	} 
}

func chanBuffer() {
	// 只能存一个数据，并且只有当该数据被取出时候才能存下一个数据。
	ch := make(chan int) // without buffer, it can't store data only when the original data was taken
	go chanBuffer_insert(ch)
	for j := 0; j < 10; j++ {
		fmt.Println(<- ch, "got from ch")
	}
}

func doTask(taskChan chan int, resultChan chan int, exitChan chan bool) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("doTask error:", err)
			return
		}	
	} ()

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
		close(taskChan)
	} ()
	
	// do tasks
	for i := 0; i < 5; i++ {
		go doTask(taskChan, resultChan, exitChan)
	}

	go func () {
		// what does exitChan do??
		// answer: to wait for doTask() finished
		for i := 0; i < 5; i++ {
			<- exitChan
		}
		close(resultChan)
		close(exitChan)
	} ()

	for res := range resultChan {
		fmt.Println(res, "is Done")
	}
}

func main() {
	// define()
	
	// operation()

	// chanBuffer()

	taskPool()	
}



















