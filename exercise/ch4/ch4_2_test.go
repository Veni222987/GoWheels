package ch4

import (
	"fmt"
	"testing"
)

func TestIsPal(t *testing.T) {
	str := "N上山到1山上N"
	fmt.Printf("len of str %d\n", len(str))
	fmt.Println(isPalindrome(str))
}
