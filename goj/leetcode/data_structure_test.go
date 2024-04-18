package leetcode

import (
	"fmt"
	"testing"
)

func TestDataStructure(t *testing.T) {
	root := buildTree("[1,2,3,null,null,4,5,null,null,null,null]")
	fmt.Printf("%d \n", root.Val)
	fmt.Printf(root.serialize())
}
