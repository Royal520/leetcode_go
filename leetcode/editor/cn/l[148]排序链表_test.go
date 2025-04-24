package main

import (
	"fmt"
	"testing"
)

func Test_SortList(t *testing.T) {
	fmt.Println("lovejj")
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func findMid(root *ListNode) *ListNode {
	fast := root.Next
	slow := root
	for fast != nil {
		if fast.Next == nil {
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	res := slow.Next
	//左右断开
	slow.Next = nil
	return res
}

func merge(left, right *ListNode) *ListNode {
	//	leetcode 21 双指针
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	h := &ListNode{}
	res := h
	for left != nil && right != nil {
		if left.Val < right.Val {
			h.Next = left
			left = left.Next
		} else {
			h.Next = right
			right = right.Next
		}
		h = h.Next
	}
	if left == nil && right != nil {
		h.Next = right
	} else if right == nil && left != nil {
		h.Next = left
	}
	return res.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := findMid(head)
	left := sortList(head)
	right := sortList(mid)
	return merge(left, right)
}

//leetcode submit region end(Prohibit modification and deletion)
