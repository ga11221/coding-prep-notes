//go:build ignore

package main

func reorderListAlt(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// 1. find middle
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 2. reverse second half
	var prev *ListNode
	curr := slow
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	second := prev

	// 3. merge: interleave first and reversed second halves
	first := head
	for second.Next != nil {
		tmp1 := first.Next
		tmp2 := second.Next
		first.Next = second
		second.Next = tmp1
		first = tmp1
		second = tmp2
	}
}
