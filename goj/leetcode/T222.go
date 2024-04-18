package leetcode

func countNodes(root *TreeNode) int {
	// 层序遍历求完全二叉树的节点个数
	if root == nil {
		return 0
	}
	var count int
	var queue = []*TreeNode{root}
	for len(queue) != 0 {
		count++
		root = queue[0]
		queue = queue[1:]
		if root.Left != nil {
			queue = append(queue, root.Left)
		}
		if root.Right != nil {
			queue = append(queue, root.Right)
		}
	}
	return count
}
