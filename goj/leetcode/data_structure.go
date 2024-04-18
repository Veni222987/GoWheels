package leetcode

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode 297 改写

// serialize 层序序列化，格式为[1,2,null,3]形式的完全二叉树
func (this *TreeNode) serialize() string {
	// []rune
	var buf bytes.Buffer
	buf.WriteRune('[')
	queue := []*TreeNode{this}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			buf.WriteString("null")
		} else {
			buf.WriteString(strconv.Itoa(node.Val))
			queue = append(queue, node.Left, node.Right)
		}
		if len(queue) != 0 {
			buf.WriteRune(',')
		}
	}
	buf.WriteRune(']')
	return buf.String()
}

// deserialize 反序列化，完全二叉树形式
func buildTree(data string) (root *TreeNode) {
	strArr := strings.Split(data[1:len(data)-1], ",")
	if len(strArr) == 0 {
		return nil
	}
	num, err := strconv.Atoi(strArr[0])
	if err != nil {
		fmt.Println("error, there are NaNs in the array or root is nil")
		return nil
	}
	root = &TreeNode{num, nil, nil}
	// 每次取两个元素
	queue := []*TreeNode{root}
	for i := 1; i < len(strArr); i += 2 {
		node := queue[0]
		queue = queue[1:]
		if strArr[i] == "null" {
			node.Left = nil
		} else {
			num, err = strconv.Atoi(strArr[i])
			if err != nil {
				fmt.Println("error, there are NaNs in the array or root is nil")
				return nil
			}
			node.Left = &TreeNode{num, nil, nil}
			queue = append(queue, node.Left)
		}

		// 处理右节点
		if i+1 >= len(strArr) {
			break
		}
		if strArr[i+1] == "null" {
			node.Right = nil
		} else {
			num, err = strconv.Atoi(strArr[i+1])
			if err != nil {
				fmt.Println("error, there are NaNs in the array or root is nil")
				return nil
			}
			node.Right = &TreeNode{num, nil, nil}
			queue = append(queue, node.Right)
		}
	}
	return root
}
