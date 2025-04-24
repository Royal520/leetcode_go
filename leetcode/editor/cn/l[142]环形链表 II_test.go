package cn

import (
	"testing"
)

func Test_LinkedListCycleIi(t *testing.T) {
	node := newListNode([]int{1, 2})
	node.Next.Next = node
	detectCycle(node)
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	nodeMap := make(map[*ListNode]int, 0)
	p := head
	for p.Next != nil {
		_, ok := nodeMap[p]
		if ok {
			return p
		}
		nodeMap[p] = p.Val
		p = p.Next
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
