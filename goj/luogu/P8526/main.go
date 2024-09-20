package main

import (
	"fmt"
	"slices"
)

// 三角dp
func main() {
	var n, m int
	fmt.Scan(&n, &m)
	list := make([]int, n)
	for i := range list {
		fmt.Scan(&list[i])
	}
	for i := 0; i < m; i++ {
		var s, t int
		fmt.Scan(&s, &t)
		num := list[s+1]
		for s = s + 2; s <= t; s++ {
			fmt.Println("list[s]=", list[s])
			num = num ^ list[s]
			fmt.Println("num", num)
		}
		fmt.Println(num)
		sum := 0
		for num != 0 {
			sum += num % 2
			num = num >> 1
		}
		fmt.Println(sum)
		slices.Sort(list)
	}

}
