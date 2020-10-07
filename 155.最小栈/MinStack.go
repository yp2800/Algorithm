package leetcode

type MinStack struct {
	Arr    []int
	Helper []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.Arr = append(this.Arr, x)
	if len(this.Helper) == 0 || x <= this.Helper[len(this.Helper)-1] {
		this.Helper = append(this.Helper, x)
	}
}

func (this *MinStack) Pop() {
	top := this.Arr[len(this.Arr)-1]
	this.Arr = this.Arr[:len(this.Arr)-1]
	if len(this.Helper) > 0 && top == this.Helper[len(this.Helper)-1] {
		this.Helper = this.Helper[:len(this.Helper)-1]
	}
}

func (this *MinStack) Top() int {
	return this.Arr[len(this.Arr)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.Helper) > 0 {
		return this.Helper[len(this.Helper)-1]
	}
	return -1
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
