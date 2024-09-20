package baiduround2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func layerTraverse(root *TreeNode) [][]int {
	queue := []*TreeNode{root}
	res := make([][]int, 0)
	for len(queue) != 0 {
		l := len(queue)
		layer := []int{}
		for i := 0; i < l; i++ {
			now := queue[i]
			layer = append(layer, now.Val)
			if now.Left != nil {
				queue = append(queue, now.Left)
			}
			if now.Right != nil {
				queue = append(queue, now.Right)
			}
		}
		queue = queue[l:]
		res = append(res, layer)
	}
	return res
}
