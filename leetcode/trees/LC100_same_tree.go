package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// So given two root nodes for two binary trees, that can have 0-2 children per node,
// walk both trees (sequentially or in parallel) and determine if they have the same
// structure and values at each node

// Edge cases: if one or both roots are nil
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 1. check for empty roots first
	// if performing in-order traversal on both trees in parallel,
	// 2. descend to left-most leaf of both trees and compare values
	// 3. if on the way there, only one of the two trees is missing a node (ie nil),
	// 		return early with false
	// 4. else return true and
	// 5. compare "this" node, and
	// 6. repeat for its right child
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	isSameLeft := isSameTree(p.Left, q.Left)
	if !isSameLeft {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Right, q.Right)
}

func main() {}

/**
    p               q
	1			    1
   / \             /  \
  2   3           2    3


    p               q
	1			    1
   /                  \
  2                    2
*/
