package main

import (
	"bufio"
	"container/heap"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

var wordRegExp = regexp.MustCompile(`\pL+('\pL+)*`)

// 从文本里统计出现的高频词汇，并打出top 10
func wordCount() {
	open, err := os.Open("./words.txt")
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(open)
	m := make(map[string]int)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatalln(err)
		}
		line := strings.ToLower(scanner.Text())
		words := wordRegExp.FindAllString(line, -1)
		for _, word := range words {
			m[word]++
		}
	}
	// 第一种思路 sort.interface
	resultWords := sortMap(m)
	for i := 0; i < 10; i++ {
		log.Println(resultWords[i])
	}
	log.Println()
	// 第二种思路 sort.slice
	resultWords = sortMap2(m)
	for i := 0; i < 10; i++ {
		log.Println(resultWords[i])
	}
	log.Println()
	// 第三种思路 heap
	stringHeapVal := sortMap3(m, 10)
	for i := 0; i < 10; i++ {
		log.Println(stringHeapVal[i])
	}
	log.Println()

	// 第四种思路 heap
	stringBigHeapVal := sortMap4(m, 10)
	for i := 0; i < 10; i++ {
		log.Println(stringBigHeapVal[i])
	}
	log.Println()

	// 第五种思路
	resultWords5 := sortMap5(m, 10)
	for i := 0; i < 10; i++ {
		log.Println(resultWords5[i])
	}
	log.Println()
}

// 第一种思路 利用sort.Interface 接口实现对数据集合的排序，实现三个方法 Len Less Swap
type wordStruct struct {
	name  string
	count int
}
type wordStructs []wordStruct

func (w wordStructs) Len() int {
	return len(w)
}
func (w wordStructs) Less(i, j int) bool {
	if w[i].count == w[j].count {
		return w[i].name > w[j].name
	}
	return w[i].count > w[j].count
}
func (w wordStructs) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func sortMap(m map[string]int) wordStructs {
	var wordsort wordStructs
	for k, v := range m {
		wordsort = append(wordsort, wordStruct{
			name:  k,
			count: v,
		})
	}
	sort.Sort(wordsort)
	return wordsort
}

// 第二种思路 利用sort.Slice 完成 []interface 的排序
func sortMap2(m map[string]int) []wordStruct {
	var wordsort []wordStruct
	for k, v := range m {
		wordsort = append(wordsort, wordStruct{name: k, count: v})
	}
	sort.Slice(wordsort, func(i, j int) bool {
		if wordsort[i].count == wordsort[j].count {
			return wordsort[i].name > wordsort[j].name
		}
		return wordsort[i].count > wordsort[j].count
	})
	return wordsort
}

// 第三种思路 小顶堆 heap top 是top k
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
		return h[i].name < h[j].name
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
	length := len(*h)
	x := (*h)[length-1]
	*h = (*h)[:length-1]
	return x
}
func sortMap3(m map[string]int, top int) stringHeap {
	wordHeap := &stringHeap{}
	for k, v := range m {
		heap.Push(wordHeap, Node{name: k, count: v})
		if wordHeap.Len() > top {
			heap.Pop(wordHeap)
		}
	}
	result := make(stringHeap, top)
	top--
	for wordHeap.Len() > 0 {
		result[top] = heap.Pop(wordHeap).(Node)
		top--
	}
	return result
}

// 第四种思路大顶堆 heap
type stringBigHeap []Node

func (h stringBigHeap) Len() int {
	return len(h)
}
func (h stringBigHeap) Less(i, j int) bool {
	//log.Println("compare... ", h[i].count, h[j].count)
	if h[i].count == h[j].count {
		return strings.Compare(h[i].name, h[j].name) > 0
	}
	return h[i].count > h[j].count
}
func (h stringBigHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *stringBigHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}
func (h *stringBigHeap) Pop() interface{} {
	//fmt.Println("xxx ", len(*h), (*h)[len(*h)-1], *h)
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
func sortMap4(m map[string]int, top int) stringHeap {
	wordHeap := &stringBigHeap{}
	for k, v := range m {
		heap.Push(wordHeap, Node{name: k, count: v}) // 这里有问题，最后一次的 push 并没有进行排序，很奇怪，而最小堆是没有问题的
		if wordHeap.Len() > top {
			//log.Println((*wordHeap)[top-1])
			heap.Pop(wordHeap)
		}
	}
	//log.Println(wordHeap)
	result := make(stringHeap, top)
	top--
	for wordHeap.Len() > 0 {
		//log.Println(wordHeap.Len(), heap.Pop(wordHeap).(Node))
		result[top] = heap.Pop(wordHeap).(Node)
		top--
	}
	return result
}
// hash 表求解

// 桶排序
func sortMap5(m map[string]int, top int) []Node {
	countWordMap := make(map[int][]string)
	for k,v := range m{
		if _,ok := countWordMap[v]; ok {
			countWordMap[v] = append(countWordMap[v], k)
			continue
		}
		countWordMap[v] = []string{k}
	}

	times := make([]int, 0, len(countWordMap))
	for k := range countWordMap {
		sort.Sort(sort.StringSlice(countWordMap[k]))
		times = append(times, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(times)))
	var nodes []Node
	for _, tim := range times {
		for _, s := range countWordMap[tim] {
			for _, s2 := range strings.Fields(s) {
				nodes = append(nodes, Node{name: s2, count: tim})
			}
		}
	}
	return nodes
}

func main() {
	wordCount()
}
