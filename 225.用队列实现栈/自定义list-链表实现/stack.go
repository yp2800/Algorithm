package leetcode

type Node struct {
	Next *Node
	Val int
}
type MyStack struct {
	top *Node
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	node := &Node{Next: this.top, Val: x}
	this.top = node
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.top == nil {return -1}
	r:=this.top.Val
	this.top.Next,this.top = nil, this.top.Next
	return r
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.top == nil {return -1}
	return this.top.Val
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.top == nil
}

