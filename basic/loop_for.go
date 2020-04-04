package main

import "fmt"

func main() {

	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	//typically "for" style in java
	for a := 0; a < 5; a++ {
		fmt.Printf("a = %d\n", a)
	}
	
	fmt.Printf("\n\n")

	//equal to while
	for a < b {
		a++
		fmt.Printf("a = %d\n", a)
	}

	fmt.Printf("\n\n")

	//a new style to iterate through the array
	for i, x := range numbers {
		fmt.Printf("No.%d is %d\n", i, x)
	}

	for x := range numbers {
		fmt.Printf("%d ", x)
	}
}

