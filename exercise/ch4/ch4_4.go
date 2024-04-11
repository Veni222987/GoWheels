package ch4

import "fmt"

/*
测试匿名结构的访问
*/

type Circle struct {
	x int
	y int
	r int
}

type Component struct {
	r int
}

type Product struct {
	Circle
	Component
	price int
}

func anonymousType() {
	p := &Product{Circle{x: 1, y: 2, r: 3}, Component{r: 4}, 5}
	fmt.Println(p.Component.r)
}
