package main

import "errors"

/*
Implement a first in first out (FIFO) queue using only two stacks.
The implemented queue should support all the functions of a normal queue (push, peek, pop, and empty).

Implement the MyQueue class:

void push(int x) Pushes element x to the back of the queue.
int pop() Removes the element from the front of the queue and returns it.
int peek() Returns the element at the front of the queue.
boolean empty() Returns true if the queue is empty, false otherwise.
Notes:

You must use only standard operations of a stack,
which means only push to top, peek/pop from top, size, and is empty operations are valid.
Depending on your language, the stack may not be supported natively.
You may simulate a stack using a list or deque (double-ended queue)
as long as you use only a stack's standard operations.


Example 1:

Input
["MyQueue", "push", "push", "peek", "pop", "empty"]
[[], [1], [2], [], [], []]
Output
[null, null, null, 1, 1, false]

Explanation
MyQueue myQueue = new MyQueue();
myQueue.push(1); // queue is: [1]
myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
myQueue.peek(); // return 1
myQueue.pop(); // return 1, queue is [2]
myQueue.empty(); // return false


Constraints:

1 <= x <= 9
At most 100 calls will be made to push, pop, peek, and empty.
All the calls to pop and peek are valid.


Follow-up: Can you implement the queue such that each operation is amortized O(1) time complexity?
In other words, performing n operations will take overall O(n) time even if one of those operations may take longer.

*/

func mainlllllll() {
	var myQ = Constructor2()
	myQ.Push(1)
	myQ.Push(2)
	println(myQ.Peek())
	println(myQ.Pop())
	println(myQ.Empty())
}

type MyStack1 struct {
	data []int
}

func (stack *MyStack1) push(x int) {
	stack.data = append([]int{x}, stack.data...)
}

func (stack *MyStack1) peek() (int, error) {
	if len(stack.data) == 0 {
		return 0, errors.New("peek on empty stack")
	}
	return stack.data[0], nil
}

func (stack *MyStack1) pop() (int, error) {
	if len(stack.data) == 0 {
		return 0, errors.New("pop on empty stack")
	}
	var d = stack.data[0]
	if len(stack.data) > 1 {
		stack.data = stack.data[1:len(stack.data)]
	} else {
		stack.data = []int{}
	}
	return d, nil
}

func (stack *MyStack1) size() int {
	return len(stack.data)
}

func (stack *MyStack1) empty() bool {
	return stack.size() == 0
}

type MyQueue struct {
	reverseOrder MyStack1
	queueOrder   MyStack1
}

func Constructor2() MyQueue {
	return MyQueue{
		reverseOrder: MyStack1{
			data: []int{},
		},
		queueOrder: MyStack1{
			data: []int{},
		},
	}
}

func (this *MyQueue) Push(x int) {
	this.reverseOrder.push(x)
}

func (this *MyQueue) Pop() int {
	if !this.queueOrder.empty() {
		pop, _ := this.queueOrder.pop()
		return pop
	}
	for this.reverseOrder.size() > 1 {
		pop, _ := this.reverseOrder.pop()
		this.queueOrder.push(pop)
	}
	ret, _ := this.reverseOrder.pop()
	return ret
}

func (this *MyQueue) Peek() int {
	if this.queueOrder.empty() {
		for this.reverseOrder.size() > 0 {
			pop, _ := this.reverseOrder.pop()
			this.queueOrder.push(pop)
		}
	}
	peek, _ := this.queueOrder.peek()
	return peek
}

func (this *MyQueue) Empty() bool {
	return this.reverseOrder.empty() && this.queueOrder.empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
