package main

func main() {
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
		f(node, k) =
			    | f(node.Left, k - node.Val)    if k - node.Val < node.Val
	        	| f(node.Right, k - node.Val)	if k - node.Val > node.Val
	        	| f(i+1, j)           			if nums[i] + nums[j] < target
*/
func findTarget(root *TreeNode, k int) bool {

}
