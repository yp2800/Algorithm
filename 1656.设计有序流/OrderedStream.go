package leetcode

import "fmt"

type OrderedStream struct {
	ptr    int
	stream []string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		ptr:    1,
		stream: make([]string, n+1),
	}
}
func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.stream[idKey] = value
	fmt.Printf("%#v\n", this.stream)
	start := this.ptr
	for this.ptr < len(this.stream) && this.stream[this.ptr] != "" {
		fmt.Println(this.ptr, "here")
		this.ptr++
	}
	println(start, this.ptr)
	return this.stream[start:this.ptr]
}
