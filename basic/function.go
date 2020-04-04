package main

import "fmt"

func max(num1, num2 int) int {
	var result int
	if num1 < num2 { 
		result = num2
	} else { 
		result = num1 
	}
	return result
}

//go can return serveral values
func swap(str1, str2 string) (string, string) {
	return str2, str1
}

func main() {
	var a int = 1
	var b int = 2
	var ret = max(a, b)

	fmt.Printf("the bigger is %d\n", ret)

	//be care that multi returned values
	str1, str2 := swap("good", "bad")

	fmt.Println(str1, str2)
}


