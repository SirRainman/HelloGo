package main

import "fmt"

func main() {
	//define a array with values
	var a = [3] int {1, 2, 3}
	
	for i, x := range a {
		fmt.Printf("No.%d is %d\n", i, x)
	} 

	//define a array without values
	var b [3] int

	//look at the way "len()" to get array's length
	fmt.Println("length of b is ",len(b))
	
	var c = [3] int {100, 200, 300}
	var ptr * int
	
	for i := 0; i < 3; i++ {
		ptr = &c[i]
		fmt.Printf("the address of c[%d] is %d, its value is %d\n", i, ptr, *ptr) 
	}

}

