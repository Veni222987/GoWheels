package main

import (
	"fmt"
)

func testCase1() {
	m := Constructor()
	m.Push(-2)
	m.Push(0)
	m.Push(-3)
	fmt.Println(m.GetMin())
	m.Pop()
	fmt.Println(m.Top())
	fmt.Println(m.GetMin())
}
