package main

import (
	"fmt"
	"math"
)

func main() {
	dp := make([]int, 1e6+1)
	var n int
	fmt.Scan(&n)
	dp[1], dp[5], dp[11] = 1, 1, 1
	for i := 1; i <= n; i++ {
		if i != 1 && i != 5 && i != 11 {
			a1, a2, a3 := math.MaxInt, math.MaxInt, math.MaxInt
			if i-1 >= 0 {
				a1 = dp[i-1]
			}
			if i-5 >= 0 {
				a2 = dp[i-5]
			}
			if i-11 >= 0 {
				a3 = dp[i-11]
			}
			dp[i] = min(a1, a2, a3) + 1
		}
	}
	fmt.Println(dp[n])
}
