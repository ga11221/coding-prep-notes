package main

type StackNode struct {
	val  int
	next *StackNode
}

type MinStack struct {
	//Doesn't necessarily have to be a stack or a slice that gets prepended, could have been appended in O(1) time ie treat end of list as head
	stackHead    *StackNode
	minStackHead *StackNode
}

/**
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2
*/

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(value int) {
	// track min
	if this.minStackHead == nil || value <= this.minStackHead.val {
		newNode := StackNode{
			val: value,
		}
		oldMinHead := this.minStackHead
		this.minStackHead = &newNode
		newNode.next = oldMinHead
	}
	newNode := StackNode{
		val: value,
	}

	// push onto actual stack
	oldHead := this.stackHead
	newNode.next = oldHead
	this.stackHead = &newNode
}

func (this *MinStack) Pop() {
	oldHead := this.stackHead
	newHead := this.stackHead.next
	this.stackHead = newHead

	// update minStack if necessary
	if this.minStackHead.val == oldHead.val {
		oldMin := this.minStackHead
		this.minStackHead = oldMin.next
		oldMin.next = nil
	}
}

func (this *MinStack) Top() int {
	return this.stackHead.val
}

func (this *MinStack) GetMin() int {
	return this.minStackHead.val
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(value);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
