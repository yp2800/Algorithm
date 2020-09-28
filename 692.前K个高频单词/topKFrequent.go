package leetcode

import (
	"container/heap"
)

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	h := &stringHeap{}
	for key, v := range m {
		heap.Push(h, Node{name: key, count: v})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([]string, k)
	var i = k - 1
	for h.Len() > 0 {
		ans[i] = heap.Pop(h).(Node).name
		i--
	}
	return ans
}

type Node struct {
	name  string
	count int
}
type stringHeap []Node

func (h stringHeap) Len() int {
	return len(h)
}
func (h stringHeap) Less(i, j int) bool {
	if h[i].count == h[j].count {
		return h[i].name > h[j].name
	}
	return h[i].count < h[j].count
}
func (h stringHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *stringHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}
func (h *stringHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
