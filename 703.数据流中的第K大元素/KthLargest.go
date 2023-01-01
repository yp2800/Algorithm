package leetcode_703

import (
	"container/heap"
	"sort"
)

type KthLargest struct {
	sort.IntSlice
	k int
}

func (kl *KthLargest) Pop() interface{} {
	//TODO implement me
	a := kl.IntSlice
	v := a[len(a)-1]
	kl.IntSlice = a[:len(a)-1]
	return v
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{
		k: k,
	}
	for _, num := range nums {
		kl.Add(num)
	}
	return kl
}

func (kl *KthLargest) Push(v interface{}) {
	kl.IntSlice = append(kl.IntSlice, v.(int))
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(kl, val)
	if kl.Len() > kl.k {
		heap.Pop(kl)
	}
	return kl.IntSlice[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
