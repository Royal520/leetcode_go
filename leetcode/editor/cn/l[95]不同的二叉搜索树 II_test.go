package cn

import (
	"fmt"
	"testing"
)

func Test_UniqueBinarySearchTreesIi(t *testing.T) {
	arry := []int{0, 1, 0, 3, 0, 0, 0, 5}
	generateTrees(1)
	root := TreeNodeLevelOrder(arry, 1)
	fmt.Println(root.Val)
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

func generateTrees(n int) []*TreeNode {
	treeArray := make([]*TreeNode, 0)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	var dfs = func(node *TreeNode, size int) {
		if size == n {
			treeArray = append(treeArray, node)
		}
		for i := 0; i < n; i++ {
			if nums[i] > 0 {

			}
		}
	}
	for i := 0; i < n; i++ {
		nums[i] = -nums[i]
		dfs(&TreeNode{Val: i + 1}, 1)
		nums[i] = -nums[i]
	}

	return treeArray
}

//leetcode submit region end(Prohibit modification and deletion)
