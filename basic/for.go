package main

import (
	"fmt"
	"math"
)

func usage_for() {
	var b int = 15
	var a int

	//typically "for" style in java
	for a := 0; a < 5; a++ {
		fmt.Printf("a = %d, ", a)
	}

	fmt.Printf("\n\n")

	//equal to while
	for a < b {
		a++
		fmt.Printf("a = %d, ", a)
	}

	fmt.Printf("\n\n")

	numbers := [6]int{1, 2, 3, 5}
	//a new style to iterate through the array
	for i, x := range numbers {
		fmt.Printf("No.%d is %d\n", i, x)
	}
	// just index
	for i := range numbers {
		fmt.Printf("%d ", i)
	}
	// just value
	for _, value := range numbers {
		fmt.Printf("%d ", value)
	}
}

func sqrt(x float64) float64 {
	res := float64(1)
	const accuracy float64 = 1e-10
	for last := float64(0); math.Abs(last-res) > accuracy; res = res - (res*res-x)/(2*res) {
		last = res
	}
	return res
}

func main() {
	// usage_for()
	fmt.Println(
		sqrt(2),
		math.Sqrt(2),
	)
}
