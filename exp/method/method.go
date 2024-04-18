package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) distanceTo(q Point) float64 {
	return math.Sqrt((p.X-q.X)*(p.X-q.X) + (p.Y-q.Y)*(p.Y-q.Y))
}

type Point2 struct {
	X, Y float64
}

func (p *Point2) scaleBy(size float64) {
	p.X, p.Y = size*p.X, size*p.Y
}

func (p *Point2) distanceTo(q *Point2) float64 {
	return math.Sqrt(((*p).X-(*q).X)*((*p).X-(*q).X) + (p.Y-q.Y)*(p.Y-q.Y))
}

func (p *Point2) length() float64 {
	if p == nil {
		return 0
	}
	return p.X*p.X + p.Y*p.Y
}

// 指针接收者方法
//func (p *Point) distance2(q *Point) float64 {
//
//}

// 报错，不允许同名
//func (p Point) X() int {
//	return 0
//}

// 基本类型增加方法
type age int

func (a age) olderThan(b age) bool {
	return a > b
}

type myPoint []float64

func (p myPoint) distanceTo(q myPoint) float64 {
	return math.Sqrt((p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1]))
}

//type myError error
//
//func (e myError) function() {
//
//}

//type pInt *int
//func (p pInt) function(){
//
//}

func main() {
	var np *Point2
	if np == nil {
		fmt.Println("p is nil")
	}
	fmt.Printf("%f\n", np.length())

	p := Point{1, 2} // 变量值
	fmt.Printf("%.2f\n", p.distanceTo(Point{2, 3}))
	fmt.Printf("%.2f\n", (&p).distanceTo(Point{2, 3}))

	p2 := &Point2{1, 2} // 变量指针
	p2.scaleBy(2)
	fmt.Printf("%.2f\n", p2.distanceTo(&Point2{0, 0}))
	fmt.Printf("%.2f\n", (*p2).distanceTo(&Point2{0, 0}))

	(*p2).scaleBy(0.5)
	fmt.Printf("%.2f\n", p2.distanceTo(&Point2{0, 0}))
	fmt.Printf("%.2f\n", (*p2).distanceTo(&Point2{0, 0}))
}
