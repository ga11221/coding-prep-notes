package main

import "testing"

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{
			name: "1->1->2->3->3",
			head: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		},
		{
			name: "empty list",
			head: nil,
			want: nil,
		},
		{
			name: "single element",
			head: &ListNode{Val: 1},
			want: &ListNode{Val: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteDuplicates(tt.head)
			for g, w := got, tt.want; g != nil && w != nil; g, w = g.Next, w.Next {
				if g.Val != w.Val {
					t.Errorf("deleteDuplicates() = %v, want %v", listToSlice(got), listToSlice(tt.want))
				}
			}
		})
	}
}

func listToSlice(head *ListNode) []int {
	var result []int
	for curr := head; curr != nil; curr = curr.Next {
		result = append(result, curr.Val)
	}
	return result
}
