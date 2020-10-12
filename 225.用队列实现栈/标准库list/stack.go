package leetcode

import "container/list"

type MyStack struct {
	*list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{list.New()}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.PushBack(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.Len() == 0 {return -1}
	return this.Remove(this.Back()).(int)
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.Len()==0{return -1}
	return this.Back().Value.(int)
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.Len() == 0
}
