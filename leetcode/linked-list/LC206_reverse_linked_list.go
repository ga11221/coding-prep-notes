package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// [2,3,4,5, 6]

func main() {
	ints := []int{2, 3, 4, 5, 6}
	head := &ListNode{
		Val: ints[0],
	}
	curr := head
	for _, v := range ints[1:] {
		node := &ListNode{
			Val: v,
		}
		curr.Next = node
		curr = node
	}

	printList(head)

	println()

	printList(reverseList(head))

}

func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("%v -> ", curr.Val)
		curr = curr.Next
	}
}

func reverseList(head *ListNode) *ListNode {
	// curr = 2
	// s:= save curr Next s:= 3
	// n:=save what 3 points to (4)
	// repoint 3 to 2   3->2 curr.Next = curr
	// advance to s curr = s

	// curr = 3
	// s:=save currents Next so that we may advance to it ( n from previous iter) (4) s:= n (4)
	// n:=save what 4 points to (5) n:= 5
	// repoint saved Next (4) to curr 4-> 3 // 4.Next = curr
	// advance to 4 (s)

	// curr = 4
	// s: save current's Next (n) s = 5
	// save what 5 points to (6) n = s.Next
	// repoint saved Next (5) to curr 5->4 s.Next = curr
	// advance to 5 curr= s

	curr := head                 // 2
	nextCurr := curr.Next        // 3
	nextOfNext := curr.Next.Next // 4
	curr.Next.Next = curr        // 3->2
	curr.Next = nil
	curr = nextCurr         // 3
	for nextOfNext != nil { // 4    // 5       // 6               // nil
		nextCurr := nextOfNext       // 4    // 5       // 6
		nextOfNext = nextOfNext.Next // 5    // 6       // nil
		nextCurr.Next = curr         // 4->3 // 5-> 4   // 6->5
		curr = nextCurr              // 4   // 5     // 6
	}
	return curr
}

func reverseList_alt(head *ListNode) *ListNode {
	var prev *ListNode
	for curr := head; curr != nil; {
		curr.Next, prev, curr = prev, curr, curr.Next
		// 2 -> nil, prev = 2, curr = 3
		// 3 -> 2, prev = 3, curr = 4
		// 4 -> 3, prev = 4, curr = 5
	}
	// [2,3,4,5]
	/*
	  			prev = nil, curr = 2
		2->nil, prev = 2 , curr =3
		3 -> 2, prev =3, curr=4
	*/
	prev := nil
	curr := head
	for true {
		next = curr.Next // 3 		 // 4
		curr.Next = prev // 2 -> nil // 3 -> 2
		prev = curr      // 2        // 3
		curr = next      // 3        // 4
	}
}
