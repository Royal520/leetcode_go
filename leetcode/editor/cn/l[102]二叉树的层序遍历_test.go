package cn

import (
	"testing"
)

//type TreeNode util.TreeNode

func Test_BinaryTreeLevelOrderTraversal(t *testing.T) {
	var order *TreeNode = TreeNodeLevelOrder([]int{0, 1, 2, 3, 4, 5}, 1)
	levelOrder(order)
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int(nil)
	}

	res := [][]int{[]int(nil)}
	cur := []*TreeNode{root}
	curIdx := 0
	// 存储当前行的终点
	end := 0

	for i := 0; i < len(cur); i++ {
		res[curIdx] = append(res[curIdx], cur[i].Val)

		if cur[i].Left != nil {
			cur = append(cur, (*TreeNode)(cur[i].Left))
		}
		if cur[i].Right != nil {
			cur = append(cur, (*TreeNode)(cur[i].Right))
		}

		if i == end && len(cur)-1 > end {
			res = append(res, []int(nil))
			curIdx++
			end = len(cur) - 1
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
