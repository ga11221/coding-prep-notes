package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

/*
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var head *ListNode
	if list1.Val <= list2.Val {
		head = list1
	} else {
		head = list2
	}
	var prev *ListNode
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			for list1 != nil && list1.Val <= list2.Val {
				prev = list1
				list1 = list1.Next
			}
			prev.Next = list2
		} else {
			for list2 != nil && list2.Val < list1.Val {
				prev = list2
				list2 = list2.Next
			}
			prev.Next = list1
		}
	}
	if list1 == nil {
		prev.Next = list2
	} else if list2 == nil {
		prev.Next = list1
	}
	return head

}
