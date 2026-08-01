package main

/*
You are given a nested list of integers nestedList. Each element is either an integer or a list whose elements may also be integers or other lists. Implement an iterator to flatten it.

Implement the NestedIterator class:

NestedIterator(List<NestedInteger> nestedList) Initializes the iterator with the nested list nestedList.
int next() Returns the next integer in the nested list.
boolean hasNext() Returns true if there are still some integers in the nested list and false otherwise.
Your code will be tested with the following pseudocode:

initialize iterator with nestedList
res = []
while iterator.hasNext()
    append iterator.next() to the end of res
return res
If res matches the expected flattened list, then your code will be judged as correct.


Example 1:

Input: nestedList = [[1,1],2,[1,1]]
Output: [1,1,2,1,1]
Explanation: By calling next repeatedly until hasNext returns false, the order of elements returned by next should be: [1,1,2,1,1].
Example 2:

Input: nestedList = [1,[4,[6]]]
Output: [1,4,6]
Explanation: By calling next repeatedly until hasNext returns false, the order of elements returned by next should be: [1,4,6].


Constraints:

1 <= nestedList.length <= 500
The values of the integers in the nested list is in the range [-106, 106].
*/

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

// DO NOT SUBMIT NESTEDINTEGER IMPL - WILL BE PROVIDED AT RUNTIME
type NestedInteger struct{}

func (this NestedInteger) IsInteger() bool {
	return false
}
func (this NestedInteger) GetInteger() int {
	return 0
}
func (n *NestedInteger) SetInteger(value int)      {}
func (this *NestedInteger) Add(elem NestedInteger) {}
func (this NestedInteger) GetList() []*NestedInteger {
	return nil
}

/*--------NestedInteger End--------------*/

type NestedPointer struct {
	ptrStack  []int
	listStack [][]*NestedInteger
}
type NestedIterator struct {
	nestedPointer NestedPointer
	nestedList    []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	pointer := NestedPointer{
		ptrStack:  []int{-1},
		listStack: [][]*NestedInteger{nestedList},
	}
	return &NestedIterator{nestedPointer: pointer, nestedList: nestedList}
}

func (this *NestedIterator) CurrPtr() int {
	return this.nestedPointer.ptrStack[len(this.nestedPointer.ptrStack)-1]
}

func (this *NestedIterator) IncrementPtr() {
	this.nestedPointer.ptrStack[len(this.nestedPointer.ptrStack)-1] =
		this.nestedPointer.ptrStack[len(this.nestedPointer.ptrStack)-1] + 1
}

func (this *NestedIterator) CurrList() []*NestedInteger {
	return this.nestedPointer.listStack[len(this.nestedPointer.listStack)-1]
}

func (this *NestedIterator) PopList() {
	this.nestedPointer.listStack = this.nestedPointer.listStack[0 : len(this.nestedPointer.listStack)-1]
	this.nestedPointer.ptrStack = this.nestedPointer.ptrStack[0 : len(this.nestedPointer.ptrStack)-1]
}

func (this *NestedIterator) IsLastList() bool {
	return len(this.nestedPointer.ptrStack) == 1
}

func (this *NestedIterator) PushList(nestedList []*NestedInteger) {
	this.nestedPointer.listStack = append(this.nestedPointer.listStack, nestedList)
	this.nestedPointer.ptrStack = append(this.nestedPointer.ptrStack, -1)
}

func (this *NestedIterator) Next() int {
	var ptr = this.CurrPtr()
	nestedInteger := this.CurrList()[ptr]
	return nestedInteger.GetInteger()
}

func (this *NestedIterator) HasNext() bool {
	/*
		listStack = [[1,1],2,[1,1]]
		  ptrStack = [-1]
		  1. increment ptrStack[last]
		  2. is ptrStack[last] == len(listStack[last]) ->
				yes,
					is len(ptrStack) == 1, return false,
					else remove last element of ptrStack and last list of listStack and goto 1.
		  3. is ptrStack[last] integer, return true
		  4. is ptrStack[last] list -> increment ptrStack[last] and append -1 to ptrStack[0, -1],
				add list = [] to listStack and repeat 1 and 2
	*/
	this.IncrementPtr()           // [1, 2]
	currPtr := this.CurrPtr()     // 0, 0, 1, 2
	currList := this.CurrList()   // [[1,1],2,[1,1]], [1,1], [1,1], [1,1]
	if currPtr == len(currList) { // no, no, no, yes
		if this.IsLastList() { // no
			return false
		}
		this.PopList() // listStack->[[1,1],2,[1,1]] ptrStack [1]
		return this.HasNext()
	}
	nestedInteger := currList[currPtr] // [1,1], 1, 1
	if nestedInteger.IsInteger() {
		return true
	}
	this.PushList(nestedInteger.GetList()) // listStack->[[1,1],2,[1,1]], [1,1] ptrStack [1, -1]
	return this.HasNext()

}
