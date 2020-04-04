package main 

import "fmt" 

func main() {
	var x interface{}
	switch i := x.(type) {
		case nil:  
			fmt.Printf(" x 的类型 :%T",i)                
		case int:  
			fmt.Printf("x 是 int 型")                      
		case float64:
			fmt.Printf("x 是 float64 型")          
		case func(int) float64:
			fmt.Printf("x 是 func(int) 型")                      
		case bool, string:
			fmt.Printf("x 是 bool 或 string 型" )      
		default:
			fmt.Printf("未知型")  
	}

	//注意fallthrough的用法
	//switch语句一般会默认最后一句话是带有break，但使用fallthrough之后，即使下一个case是false，也会强制执行下一个case

	switch {
		case false:
			fmt.Println("1、case 条件语句为 false")
			fallthrough
		case true:
			fmt.Println("2、case 条件语句为 true")
			fallthrough
		case false:
			fmt.Println("3、case 条件语句为 false")
			fallthrough
		case true:
			fmt.Println("4、case 条件语句为 true")
		case false:
			fmt.Println("5、case 条件语句为 false")
			fallthrough
		default:
			fmt.Println("6、默认 case")
    }
}
