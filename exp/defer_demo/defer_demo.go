package main

import "fmt"

func plus(x int) int {
	defer func() {
		x++
	}()
	return x * 2
}

func plus2(x int) (res int) {
	res = x
	defer func() {
		res++
	}()
	return res * 2
}

func plus3(x *int) *int {
	defer func() {
		*x++
	}()
	*x = *x * 2
	return x
}

func main() {
	x := 10
	myfunc := func() {
		fmt.Println("x的值为:", x)
		x += 5
		fmt.Println("修改后:", x)
	}
	myfunc()
	myfunc()
	fmt.Println("匿名函数执行后:", x)
}
