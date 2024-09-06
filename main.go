package main

import "fmt"

func main() {
	str := "abc"
	//arr := []int{5, 6, 7, 8}
	// !增强型for循环遍历
	for _, v := range str {
		fmt.Println(v)
	}

	// !增强型for循环遍历数组
	for v := range 10 {
		fmt.Println(v)
	}
}
