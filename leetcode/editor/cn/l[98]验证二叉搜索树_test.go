package cn

import (
	"fmt"
	"math"
	"testing"
)

func Test_ValidateBinarySearchTree(t *testing.T) {
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
var base int64

func isValidBST(root *TreeNode) bool {
	base = math.MinInt64
	return inorderTraversal(root)
}

func inorderTraversal(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if inorderTraversal((*TreeNode)(root.Left)) != true {
		return false
	}
	if int64(root.Val) <= base {
		return false
	}
	base = int64(root.Val)
	if inorderTraversal((*TreeNode)(root.Right)) != true {
		return false
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
