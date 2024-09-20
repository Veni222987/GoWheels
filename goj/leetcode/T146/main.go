package main

func main() {
	l := Constructor(2)
	l.Put(1, 1)
	l.Put(2, 2)
	l.Get(1)
	l.Put(3, 3)
	l.Get(2)
	l.Put(4, 4)
	l.Get(1)
	l.Get(3)
	l.Get(4)
}
