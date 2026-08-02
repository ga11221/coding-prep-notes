package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.

f(nil, _) = 0,false
f(root, c) = c<-f(left, c) +1 if c == k return left f(right, c)

*/

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}
	println(kthSmallest(root, 4))
}

func kthSmallest(root *TreeNode, k int) int {
	val, _ := _kthSmallest(root, k, 0)
	return val
}

func _kthSmallest(root *TreeNode, k int, c int) (int, bool) {
	if root == nil {
		return c, false
	}
	val, found := _kthSmallest(root.Left, k, c)
	if found {
		return val, true
	}
	val++
	if val == k {
		return root.Val, true
	}
	return _kthSmallest(root.Right, k, val)
}
