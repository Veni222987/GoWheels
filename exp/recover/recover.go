package main

import "fmt"

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("捕获到panic:", r)
	}
}

func doSomething() {
	// 产生panic
	panic("发生了错误：1")
}

func run() {
	defer recoverFromPanic()
	doSomething()
}

func main() {
	run()
	fmt.Println("程序继续执行")
}
