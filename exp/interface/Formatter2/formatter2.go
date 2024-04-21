package Formatter2

type Formatter2 interface {
	MyWrite(strs ...string) (res string, err error)

	DistinctOf2() (res string)
}
