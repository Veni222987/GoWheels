package Formatter1

type Formatter1 interface {
	MyWrite(strs ...string) (res string, err error)

	DistinctOf1() (res string)
}
