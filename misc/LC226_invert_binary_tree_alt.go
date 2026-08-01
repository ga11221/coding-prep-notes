package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil 
	}
    
	left := invertTree(root.Left) //2a: invertTree(4); 4a: invertTree(8); 8a: invertTree(nil)
	right := invertTree(root.Right)// 8a: invertTree(nil);
	root.Right = left// 8a:
	root.Left = right
	return root
}
