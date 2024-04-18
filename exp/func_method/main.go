package main

import (
	"fmt"
)

/*
// twoSum
最简单的函数
*/
func twoSum(x, y int) int {
	return x + y
}

func twoSum2(x, y int) (res int) {
	res = x + y
	return
}

func sumOfAll(nums ...int) (res int) {
	if nums == nil {
		fmt.Println("slice is nil")
	} else if len(nums) == 0 {
		fmt.Println("slice is empty")
	}
	for _, value := range nums {
		res += value
	}
	return
}

func passingFunc(data []int, callback func(int) int) {
	for _, d := range data {
		fmt.Println(callback(d))
	}
}

func numDouble(num int) int {
	return num * 2
}

/*
函数的变量传递方式
*/
func greeting(name string) {
	fmt.Printf("Hello, %s\n", name)
}

/*
函数变量

*/

/*
匿名函数
*/
func processItems(items ...string) {
	if items == nil {
		fmt.Println("切片为nil")
	} else {
		fmt.Println("切片是empty")
	}
}

func main() {
	data := make([]int, 5)
	for i, _ := range data {
		data[i] = i
	}
	passingFunc(data, func(num int) int {
		return num * 2
	})

	processItems()
}
