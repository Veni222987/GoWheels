package main

// 两个同名接口

type Formatter1 interface {
	myWrite(strs ...string) (res string, err error)
}

//type Formatter2 interface {
//	myWrite(strs ...string) (res string, err error)
//}

type StrArr []string

func (sa StrArr) myWrite(strs ...string) (res string, err error) {
	return "", nil
}
