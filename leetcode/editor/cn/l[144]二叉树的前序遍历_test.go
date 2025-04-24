package main

import (
	"fmt"
	"testing"
)

func Test_BinaryTreePreorderTraversal(t *testing.T) {
	fmt.Println("lovejj")
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	arry := []*TreeNode{root}
	res := []int{}
	res = append(res, root.Val)
	p := root
	for p.Left != nil || len(arry) > 0 {
		if p.Left != nil {
			arry = append(arry, p.Left)
			res = append(res, p.Left.Val)
			tmp := p.Left
			p.Left = nil
			p = tmp
		} else {
			p = arry[len(arry)-1]
			arry = arry[:len(arry)-1]
			if p.Right != nil {
				arry = append(arry, p.Right)
				res = append(res, p.Right.Val)
				p = p.Right
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
