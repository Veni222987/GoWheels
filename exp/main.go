package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "1,,2"
	ss := strings.Split(s, ",")
	for _, str := range ss {
		fmt.Println(str)
	}
}
