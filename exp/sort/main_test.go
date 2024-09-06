package main

import "testing"

// 基准测试函数
func BenchmarkPerformance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		performance(10000) // 在这里调用要进行基准测试的函数
	}
}
