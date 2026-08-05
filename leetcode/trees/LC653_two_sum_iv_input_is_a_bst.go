package main

func main() {
	root := &TreeNode{Val: 1}
	zero := &TreeNode{Val: 0}
	four := &TreeNode{Val: 4}
	negtwo := &TreeNode{Val: -2}
	three := &TreeNode{Val: 3}
	root.Left = zero
	root.Right = four
	zero.Left = negtwo
	four.Left = three
	println(findTarget(root, 7))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrder(root *TreeNode, smallest chan<- int) {
	if root == nil {
		return
	}
	inOrder(root.Left, smallest)
	smallest <- root.Val
	inOrder(root.Right, smallest)
}

func reverseInOrder(root *TreeNode, largest chan<- int) {
	if root == nil {
		return
	}
	reverseInOrder(root.Right, largest)
	largest <- root.Val
	reverseInOrder(root.Left, largest)
}

func findTarget(root *TreeNode, k int) bool {
	smallest := make(chan int)
	largest := make(chan int)
	go inOrder(root, smallest)
	go reverseInOrder(root, largest)
	var largerVal int
	var nextSmallest bool
next:
	for i := range smallest {
		if nextSmallest {
			if i == largerVal {
				return false
			}
			if i+largerVal < k {
				continue
			}
			if i+largerVal == k {
				return true
			}
		}
		for j := range largest {
			nextSmallest = false
			if i == j {
				return false
			}
			if i+j == k {
				return true
			}
			if i+j > k {
				continue
			}
			if i+j < k {
				largerVal = j
				nextSmallest = true
				continue next
			}
		}
	}
	return false
}
