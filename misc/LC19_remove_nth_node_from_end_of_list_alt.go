package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Given the head of a linked list, remove the nth node from the end of the list and return its head.

[1,2,3,4,5], n = 2

ret = [1,2,3,5]

linked list -> no length
must traverse starting from head

slow,fast pointers -
distance fast from slow by n
advance both slow and fast until fast = tail
repoint slow's prev to slow's Next
*/
// [1,2] n = 1
// [1] n=1
// [1,2] n=2
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// distance fast from slow by n
	slow, fast := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	} // slow = 1 fast = nil
	// INVARIANTS - slow is always n nodes away from fast, slow.Next != nil, fast.Next != nil
	if fast == nil {
		return head.Next
	}
	for fast.Next != nil {
		slow = slow.Next // 2
		fast = fast.Next // nil
	}
	// POST-INVARIANT - fast = nil, prev = nil only if slow = head (ie removing first element of linked list)
	slow.Next = slow.Next.Next
	return head
}
