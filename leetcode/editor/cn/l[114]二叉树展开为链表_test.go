package cn

import (
	"fmt"
	"testing"
)

func Test_FlattenBinaryTreeToLinkedList(t *testing.T) {
	order := TreeNodeLevelOrder([]int{0, 1, 2, 5, 3, 4, 0, 6}, 1)
	fmt.Println(order)
	beforeRead(order)
	flatten(order)
	fmt.Println(order)
}
func beforeRead(root *TreeNode) {
	arry := []*TreeNode{root}
	fmt.Println(root.Val)
	dumNode := &TreeNode{}
	dumNode.Right = root
	p := dumNode.Right
	for p.Left != nil || len(arry) > 0 {
		if p.Left != nil {
			arry = append(arry, (*TreeNode)(p.Left))
			fmt.Println(p.Left.Val)
			tmp := p.Left
			p.Left = nil
			p = tmp
		} else {
			p = arry[len(arry)-1]
			arry = arry[:len(arry)-1]
			if p.Right != nil {
				arry = append(arry, p.Right)
				fmt.Println(p.Right.Val)
				p = p.Right
			}
		}
	}
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
//

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten0(root)
}

func flatten0(root *TreeNode) (first, last *TreeNode) {
	first, last = root, root
	left, right := root.Left, root.Right
	root.Left, root.Right = nil, nil
	if left != nil {
		last.Right, last = flatten0(left)
	}
	if right != nil {
		last.Right, last = flatten0(right)
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
