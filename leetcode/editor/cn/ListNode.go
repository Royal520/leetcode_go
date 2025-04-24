package main

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
