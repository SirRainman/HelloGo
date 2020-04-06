package main

import "fmt"

// 详细资料：https://blog.golang.org/slices-intro
func main() {

	// fmt.Printf("to test append() and copy()\n")
	// sliceUsage()

	// 在这个函数中有疑问？
	// len 和 cap 的变化规律并不好发现
	fmt.Printf("\n\nto test append() and copy()\n")
	test_Append_Copy()
}

func sliceUsage() {
	/* 创建切片，注意没有指明大小 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5) // len=0 cap=5 slice=[]
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2] // len=2 cap=9 slice=[0 1]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5] // len=3 cap=7 slice=[2 3 4], 注意这里的cap发生了变化，相当于除去了0 1
	printSlice(number3)

	//在这里注意cap的变化
	//我们能注意到如果是cap的大小是与起点的位置有关的，起点越靠前，cap越大
	number4 := numbers[5:] // len=4 cap=4 slice=[5 6 7 8]
	printSlice(number4)
}

func test_Append_Copy() {
	var numbers []int // len=0 cap=0 slice=[]，此时numbers == nil
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0) // len=1 cap=1 slice=[0]
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1) // len=2 cap=2 slice=[0 1]
	printSlice(numbers)

	/* 同时添加多个元素 */
	//注意这里cap的大小开始比len大一个
	numbers = append(numbers, 2, 3, 4) // len=5 cap=6 slice=[0 1 2 3 4]
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2) // len=5 cap=12 slice=[0 0 0 0 0]
	printSlice(numbers1)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers) // len=5 cap=12 slice=[0 1 2 3 4]
	printSlice(numbers1)

	// 注意numbers的cap没有变化，只有len的值发生了变化，与前面情况不同了
	numbers = append(numbers, 5) // len=6 cap=6 slice=[0 1 2 3 4 5]
	printSlice(numbers)

	//注意cap的变化，为什么突然增加到了12？
	numbers = append(numbers, 6, 7, 8) // len=9 cap=12 slice=[0 1 2 3 4 5 6 7 8]
	printSlice(numbers)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
