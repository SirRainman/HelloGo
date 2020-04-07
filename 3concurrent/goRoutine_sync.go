package main

import (
    "fmt"
    "sync"
)

func cal(a int, b int, waitNum *sync.WaitGroup) {
	fmt.Printf("%d + %d = %d\n", a, b, a + b)
	defer waitNum.Done() //waitNum-- when goroutine finished 
}

func main() {
	var waitNum sync.WaitGroup 
	for i := 0 ; i < 10; i++ {
		waitNum.Add(1) // waitNum++
		go cal2(i, i, &waitNum)
	}
	waitNum.Wait() // wait all goRoutine finishes
}
