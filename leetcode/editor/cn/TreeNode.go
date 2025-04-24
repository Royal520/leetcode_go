package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TreeNodeLevelOrder(array []int, i int) *TreeNode {
	if i >= len(array) || array[i] == 0 {
		return nil
	}
	root := &TreeNode{Val: array[i]}
	root.Left = TreeNodeLevelOrder(array, 2*i)
	root.Right = TreeNodeLevelOrder(array, 2*i+1)
	return root

}
