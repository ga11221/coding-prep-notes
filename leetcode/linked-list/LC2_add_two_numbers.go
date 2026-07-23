package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
*

	Input: l1 = [2,4,3], l2 = [5,6,4]
	Output: [7,0,8]
	always left align lists and add carry bit to Next num
	Explanation: 342 + 465 = 807.
*/

/*
*
l1=[9,9,9,9,9,9,9]
l2=[9,9,9,9]
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	curr1 := l1 // 9
	curr2 := l2 // 9
	head := &ListNode{}
	curr3 := head

	carry := 0
	var cont *ListNode
	var prev *ListNode
	for true {
		sum := curr1.Val + curr2.Val + carry // 18 // 18 + 1 // 18 + 1 // 18+1
		if sum > 9 {
			sum %= 10 // 8 // 9 // 9 // 9
			carry = 1
		} else {
			carry = 0
		}
		curr3.Val = sum // 8 -> 9 -> 9 -> 9
		tmp := &ListNode{}
		curr3.Next = tmp
		prev = curr3
		curr3 = tmp
		if curr1.Next == nil {
			cont = curr2
			break
		} else if curr2.Next == nil {
			cont = curr1 // 9
			break
		}
		curr1 = curr1.Next // 9 // 9 // 9
		curr2 = curr2.Next // 9 // 9 // 9
	}
	for cont.Next != nil {
		cont = cont.Next        // 5th 9 // 6th 9 // 7th 9
		sum := cont.Val + carry // 10 // 10 // 10
		if sum > 9 {
			sum %= 10 // 0 // 0 // 0
			carry = 1
		} else {
			carry = 0
		}
		curr3.Val = sum // 0 // 0 // 0
		tmp := &ListNode{}
		curr3.Next = tmp
		prev = curr3 // 0
		curr3 = tmp
	}
	if carry == 1 {
		curr3.Val = 1
	} else {
		prev.Next = nil
	}

	return head
}
