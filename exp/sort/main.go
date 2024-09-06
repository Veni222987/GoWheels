package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sort"
	"time"
)

// 主要研究排序的用法和性能

func main() {
	performance(100000000)
}

// 展示如何使用slices包排序
func usage() {
	arr := []int{5, 7, 3, 8, 2, 1}
	// 默认升序排列
	slices.Sort(arr)
	fmt.Println(arr)
	// 自定义排列函数实现降序排列
	slices.SortFunc(arr, func(i, j int) int { return j - i })
	fmt.Println(arr)
	// 自定义类型排列
	type person struct {
		name string
		age  int
	}
	parr := []person{
		{"momo", 18},
		{"neo", 22},
		{"venii", 15},
	}
	slices.SortFunc(parr, func(a, b person) int { return a.age - b.age })
	fmt.Println(parr)
	// 自定义非数据类型排序，例如按照名字长度排序
	slices.SortFunc(parr, func(a, b person) int { return len(a.name) - len(b.name) })
	fmt.Println(parr)
}

// 测试性能，比较sort和slice排序速度
func performance(length int) {
	// 随机初始化数组
	arr1, arr2 := make([]int, length), make([]int, length)
	for i := range length {
		randNum := rand.Intn(101)
		arr1[i] = randNum
		arr2[i] = randNum
	}

	start := time.Now()
	sort.Slice(arr1, func(i, j int) bool { return arr1[i] < arr1[j] })
	timeMicroS := time.Since(start).Microseconds()
	fmt.Println("time of arr1:", timeMicroS)

	start = time.Now()
	slices.SortFunc(arr2, func(a, b int) int { return a - b })
	timeMicroS = time.Since(start).Microseconds()
	fmt.Println("time of arr2:", timeMicroS)
}
