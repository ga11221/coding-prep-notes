//go:build ignore

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	curr := head
	val := curr.Val
	var prev *ListNode
	for curr.Next != nil {
		prev = curr
		curr = curr.Next
		if curr.Val == val {
			prev.Next = curr.Next
			curr.Next = nil
		} else {
			val = curr.Val
		}
	}
	return head
}
