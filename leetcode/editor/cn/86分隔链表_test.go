package main

import (
	"fmt"
	"testing"
)

func Test_partition(t *testing.T) {
	fmt.Println("")
}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// leetcode submit region begin(Prohibit modification and deletion)
func partition(head *ListNode, x int) *ListNode {
	dummyHead := &ListNode{}
	cur := head
	//存储小的值
	small := dummyHead
	smallHead := small
	//存储当前的的值
	big := &ListNode{}
	//存储大的头结点
	bigHead := big
	for cur != nil {
		if cur.Val <= x {
			small.Next = cur
			small = small.Next
		} else {
			big.Next = cur
			big = big.Next
		}
		cur = cur.Next
	}
	dummyHead.Next = smallHead.Next
	small.Next = bigHead.Next
	return dummyHead.Next
}

//leetcode submit region end(Prohibit modification and deletion)
