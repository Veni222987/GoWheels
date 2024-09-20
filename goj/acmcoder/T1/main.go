package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		fmt.Println(solve(m))
	}
}

func solve(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return solve(n-1) + solve(n-2)
}
