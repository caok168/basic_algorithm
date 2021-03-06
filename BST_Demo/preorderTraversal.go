package main

/***
εεΊιε
*/

func preOrderTraversal(root *TreeNode) []int {
	var preOrder func(*TreeNode)
	res := []int{}

	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		preOrder(node.Left)
		preOrder(node.Right)
	}

	preOrder(root)

	return res
}
