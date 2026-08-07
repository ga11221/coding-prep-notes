package main

/*
LC94 Binary Tree Inorder Traversal (Stack, Tree, DFS)

Given the root of a binary tree, return the inorder traversal of its nodes'
values.

Input: root = [1,null,2,3]
Output: [1,3,2]

Input: root = []
Output: []

Input: root = [1]
Output: [1]

Constraints:
- The number of nodes in the tree is in the range [0, 100].
- -100 <= Node.val <= 100

Follow up: recursive solution is trivial; could you do it iteratively?

n nodes in tree
traversing through nodes, printing left most child, its parent, then its right child
f(nil) = return
f(root) = f(root.L); parent; f(root.R)
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// [1,null,2,3]
/*
		1
	nil    2
	     3

*/
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	vals := []int{}
	stack := []*TreeNode{}
	curr := root
	for curr != nil {
		push(&stack, curr)
		curr = curr.Left
	}
	for len(stack) > 0 {
		this := pop(&stack)
		vals = append(vals, this.Val)
		curr := this.Right
		for curr != nil {
			push(&stack, curr)
			curr = curr.Left
		}
	}
	return vals
}

func push(stack *[]*TreeNode, node *TreeNode) {
	*stack = append(*stack, node)
}

func pop(stack *[]*TreeNode) *TreeNode {
	popped := (*stack)[len(*stack)-1]
	*stack = (*stack)[0 : len(*stack)-1]
	return popped
}

/*
6-rung ladder (see ladder.md - the generalized 6-step ladder).

Correct approach: iterative inorder with an explicit stack, replacing the
call stack (the follow-up). Shape:

  stack := []*TreeNode{}
  cur := root
  for cur != nil || len(stack) > 0:
      for cur != nil:
          stack = append(stack, cur)   // descend leftmost, deferring each node
          cur = cur.Left
      cur = stack[len(stack)-1]; stack = stack[:len(stack)-1]  // pop
      result = append(result, cur.Val) // visit: left subtree fully done
      cur = cur.Right                  // move to right subtree

Rung 1 - ENUMERATE (ground truth):

	inorder = traverse(root.Left), visit root, traverse(root.Right). Visit
	every node exactly once; the order is defined by the recurrence. Output
	is O(n), so the enumeration IS the answer - no compression beats emitting
	n nodes.

Rung 2 - NAME THE OBJECT:

	the ordered visit sequence of all n nodes. Question type ALL: the answer
	is the whole sequence, so the output size fixes O(n) as a floor.

Rung 3 - COMPRESS THE SPACE (representation axis):

	the object of interest is "which nodes are pending". The call stack is
	replaced by an explicit stack that holds exactly the chain of ancestors
	whose left subtrees are exhausted but who are not yet visited. Small
	state: O(height), not O(n) memory for the pending frontier.

Rung 4 - COLLAPSE TO A STATISTIC (aggregation axis):

	aggregation is append in visit order. Each node contributes exactly one
	value, emitted when its left subtree is done - the sequence IS the
	statistic.

Rung 5 - FIX A COMPUTATION ORDER (time axis):

	exhaustion by the traversal order itself: go as far left as possible,
	then unwind. Every node is pushed once and popped once -> O(n) time,
	O(height) space. The pop order mirrors the recursion exactly.

Rung 6 - PROVE AN INVARIANT:

	invariant: just before each pop, the stack is, bottom to top, the chain of
	ancestors of cur whose left subtrees are finished but who are unvisited.
	Claim: a node is visited exactly when its entire left subtree has already
	been emitted. The descent loop guarantees the leftmost unvisited path is
	on the stack; after visiting cur we move to cur.Right, whose left subtree
	is emptied the same way. Induction on tree structure => each node visited
	once, in inorder order: left, self, right. Soundness: every emitted value
	is real (one per node). Completeness: every node is emitted (no node is
	skipped; the descent reaches every non-nil child). The explicit stack
	holds no more than height nodes, matching the recursion's work.

Complexity: O(n) time, O(height) space.
*/
