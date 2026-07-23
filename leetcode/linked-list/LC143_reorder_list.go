//go:build ignore

package main

/**
You are given the head of a singly linked-list. The list can be represented as:

L0 → L1 → … → Ln - 1 → Ln
Reorder the list to be on the following form:

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
You may not modify the values in the list's nodes. Only nodes themselves may be changed.
L0 points to last
Last points to L0.old.Next
L1 points to Last.prev

L0->L1....->Ln
Ln->Ln-1->Ln-2

L0 -> L1 -> L2
L2 -> L1 -> L0

[l0,l1,l2]
[l2,l1,l0]
l0->l2->l1


L0 -> L1 -> L2 -> L3
L3 -> L2 -> L1 -> L0

[l0,l3,l1,l2]
for each elem in list:
	append elem to slice orig
for each elem in list:
   prepend elem to slice rev
curr = orig[0]
toggle = true
for i from 0 to len/2:
	if toggle:
		curr.Next = rev[i] // orig[0].Next = rev[0] [L0, L2]
		curr = rev[i]
	else:
		curr.Next = orig[i] // L2.Next = L1 [L0, L2, L1]
		curr = orig[i]
	toggle = !toggle

curr.Next = nil

orig[0] -> rev[0] -> orig [1] -> rev[1]...
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// [1,2,3,4]

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	if head.Next == nil {
		return
	}
	curr := head
	originalOrder := []*ListNode{}
	reverseOrder := []*ListNode{}
	for curr != nil {
		originalOrder = append(originalOrder, curr)
		reverseOrder = append([]*ListNode{curr}, reverseOrder...)
		curr = curr.Next
	}
	// [1,2,3,4, 5]
	// [5, 4, 3,2,1,]
	// [1,2,5,4,3]
	// 1->5->2->4->3

	// [1,2,3,4]
	// [1,2,4,3]
	// [4,3,2,1]
	// 1->4->2->3
	curr = originalOrder[0]
	toggle := true
	j := 1
	for i := 0; i <= len(originalOrder)/2; {
		if toggle {
			curr.Next = reverseOrder[i] // i=0: 1 -> 5 i = 1: 2->4
			j++
			if j == len(originalOrder) {
				break
			}
			curr = reverseOrder[i]
			i++
		} else {
			curr.Next = originalOrder[i] // i:=1: 5 -> 2; i=2:4->3
			j++
			if j == len(originalOrder) {
				break
			}
			curr = originalOrder[i]
		}
		toggle = !toggle
	}
	curr.Next.Next = nil
}
