package main

import (
	"GoWheels/exp/interface/Formatter1"
	"GoWheels/exp/interface/Formatter2"
	"fmt"
)

//// 声明两个函数完全一样的接口
//
//type Formatter1 interface {
//	myWrite(strs ...string) (res string, err error)
//}
//
//type Formatter2 interface {
//	myWrite(strs ...string) (res string, err error)
//}

// 不同接口中含有同名函数的情况

type StrArr []string

func (sa StrArr) MyWrite(strs ...string) (res string, err error) {
	for _, str := range strs {
		res += str + ","
	}
	return
}

func (sa StrArr) DistinctOf1() (res string) {
	return "I'm 1"
}

func (sa StrArr) DistinctOf2() (res string) {
	return "I'm 2"
}

// 测试是否可以当formatter1使用
func formattest1(formatter1 Formatter1.Formatter1) {
	fmt.Println(formatter1.MyWrite([]string{"111", "222"}...))
}

// 测试是否可以当formatter2使用
func formattest2(formatter2 Formatter2.Formatter2) {
	fmt.Println(formatter2.MyWrite([]string{"222", "333"}...))
}

func main() {
	var strArr StrArr
	formattest2(strArr)
}
