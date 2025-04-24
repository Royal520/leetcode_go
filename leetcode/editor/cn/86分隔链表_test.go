package cn

import (
	"fmt"
	"testing"
)

// go test -v 86分隔链表_test.go ListNode.go
func Test_partition(t *testing.T) {
	node := newListNode([]int{2, 1})
	printNode(node)
	fmt.Println("")
	after := partition(node, 2)
	printNode(after)
}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// leetcode submit region begin(Prohibit modification and deletion)
func partition(head *ListNode, x int) *ListNode {
	cur := head
	//存储小的值
	small := &ListNode{}
	smallHead := small
	//存储当前的的值
	big := &ListNode{}
	//存储大的头结点
	bigHead := big
	for cur != nil {
		if cur.Val < x {
			small.Next = cur
			small = small.Next
		} else {
			big.Next = cur
			big = big.Next
		}
		cur = cur.Next
	}
	big.Next = nil
	small.Next = bigHead.Next
	return smallHead.Next
}

//leetcode submit region end(Prohibit modification and deletion)
