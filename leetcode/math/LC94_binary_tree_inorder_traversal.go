package main

import "fmt"

func mainiai() {
	//root := []interface{}{1, nil, 2, 3}
	root := []interface{}{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9}
	tree := createTree(root)
	fmt.Printf("%#v", tree)
	// todo write pretty printer for tree
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createTree(root []interface{}) *TreeNode {
	if len(root) == 0 || root[0] == nil {
		return nil
	}
	var nodes = []*TreeNode{
		{
			Val: root[0].(int),
		},
	}
	var rt = &TreeNode{}
	for i, j := 0, 1; ; {
		rt = nodes[i]
		i++
		if rt == nil {
			continue
		}
		if root[j] == nil {
			rt.Left = nil
		} else {
			rt.Left = &TreeNode{
				Val: root[j].(int),
			}
		}
		nodes = append(nodes, rt.Left)
		j++
		if j == len(root) {
			break
		}
		if root[j] == nil {
			rt.Right = nil
		} else {
			rt.Right = &TreeNode{
				Val: root[j].(int),
			}
		}
		nodes = append(nodes, rt.Right)
		j++
		if j == len(root) {
			break
		}

	}
	return nodes[0]
}

func inorderTraversal(root *TreeNode) []int {
	return nil
}
