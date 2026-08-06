package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func (c *Codec) serialize(root *TreeNode) string {
	nodes := []string{}
	if root == nil {
		return "nil"
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curr := queue[0]
		if curr == nil {
			nodes = append(nodes, "nil")
		} else {
			nodes = append(nodes, strconv.Itoa(curr.Val))
			queue = append(queue, []*TreeNode{curr.Left, curr.Right}...)
		}
		queue = queue[1:]
	}
	return strings.Join(nodes, ",")
}

func (c *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	if nodes[0] == "nil" {
		return nil
	}
	val, _ := strconv.Atoi(nodes[0])
	root := &TreeNode{Val: val}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 {
		parent := queue[0]
		queue = queue[1:]
		if i < len(nodes) && nodes[i] != "nil" {
			v, _ := strconv.Atoi(nodes[i])
			parent.Left = &TreeNode{Val: v}
			queue = append(queue, parent.Left)
		}
		i++
		if i < len(nodes) && nodes[i] != "nil" {
			v, _ := strconv.Atoi(nodes[i])
			parent.Right = &TreeNode{Val: v}
			queue = append(queue, parent.Right)
		}
		i++
	}
	return root
}

func Constructor() Codec {
	return Codec{}
}
