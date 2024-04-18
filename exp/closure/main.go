package main

import "fmt"

func newCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	f := newCounter()
	fmt.Println(f())
	fmt.Println(f())
	f1 := newCounter()
	fmt.Println(f1())
	fmt.Println(f1())
}
