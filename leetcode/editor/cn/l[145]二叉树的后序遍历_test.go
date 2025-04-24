package cn

import (
	"fmt"
	"testing"
)

func Test_BinaryTreePostorderTraversal(t *testing.T) {
	fmt.Println("lovejj")
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Right *TreeNode
 *     R *TreeNode
 * }
 */
func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func postorderTraversal(root *TreeNode) (res []int) {
	addPath := func(node *TreeNode) {
		resSize := len(res)
		for ; node != nil; node = node.Right {
			res = append(res, node.Val)
		}
		reverse(res[resSize:])
	}

	p1 := root
	for p1 != nil {
		if p2 := p1.Left; p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				p2.Right = p1
				p1 = p1.Left
				continue
			}
			p2.Right = nil
			addPath(p1.Left)
		}
		p1 = p1.Right
	}
	addPath(root)
	return
}

//leetcode submit region end(Prohibit modification and deletion)
