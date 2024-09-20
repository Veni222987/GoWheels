package main

import "fmt"

func main() {
	// 创建一个map
	myMap := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}

	// 打印删除前的map
	fmt.Println("Map before deletion:", myMap)

	// 删除键为"two"的键值对
	delete(myMap, 3)

	// 打印删除后的map
	fmt.Println("Map after deletion:", myMap)
}
