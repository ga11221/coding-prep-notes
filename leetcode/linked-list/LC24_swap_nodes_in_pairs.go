package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// step 0 - restate problem and confirm assumptions:
// So, I'm given a linked list, every 2 nodes need to be swapped, and the node swapped into the "second" position, should point to either:
// a. the next swapped pair (if there are two or more nodes left),
// b. the next unswapped node (if it's the last node in the list), or
// c. tail (if there are no nodes left)
// Do I have it right?

// step 1 - edge cases:
// empty list, singleton list, list having odd vs even length
// [], [1], [1,2], [1,2,3], [1,2,3,4]
func swapPairs(head *ListNode) *ListNode {
	// step 2 - approach:
	// 1->2->3->4->nil
	// **save new head for return**
	// 1. swap first pair (two pointers):
	// 		2->1 - 1 still points to 2 and 2 no longer points to next node in list
	// 2. so: store node that 2 pointed to (this is the node to advance to on next iteration)
	// 3. store node 1 (this is the node to repoint after the swap in the next iteration)
	// one pass O(n)  no extra space O(1)

	// []
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next // 2
	pointer1 := head     // 1
	pointer2 := newHead  // 2
	for pointer2 != nil {
		// step 2
		nodeToAdvance := pointer2.Next // 2.Next -> nil
		// step 3
		nodeToRepoint := pointer1 // 1
		// step 1
		pointer2.Next = pointer1 // 2.Next <- 1
		pointer1 = nodeToAdvance // nil
		if nodeToAdvance == nil {
			pointer2 = nil
		} else {
			pointer2 = nodeToAdvance.Next
		}
		if pointer2 == nil /*odd length - encountered last node*/ {
			nodeToRepoint.Next = pointer1 // 1.Next <- nil
		} else {
			nodeToRepoint.Next = pointer2
		}

		// after 1st iteration: [2->1->3]

	}
	return newHead

}
