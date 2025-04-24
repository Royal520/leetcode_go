package cn

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNode(arry []int) *ListNode {
	if len(arry) == 0 {
		return nil
	}
	root := &ListNode{Val: arry[0]}
	p := root
	for _, v := range arry[1:] {
		p.Next = &ListNode{Val: v}
		p = p.Next
	}
	return root
}

func printNode(node *ListNode) {
	cur := node
	for cur != nil {
		fmt.Print(cur.Val)
		fmt.Print("-->")
		cur = cur.Next
	}
}
