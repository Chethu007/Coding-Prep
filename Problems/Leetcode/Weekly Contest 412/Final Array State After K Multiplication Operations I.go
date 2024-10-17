package main

import (
	"container/heap"
	"fmt"
)

type item struct {
	val   int
	index int
}

type MinHeap []item

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	if h[i].val == h[j].val {
		return h[i].index < h[j].index
	}
	return h[i].val < h[j].val
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(item))
}

func (h *MinHeap) Pop() interface{} {
	temp := *h
	n := len(temp)
	*h = temp[:n-1]
	val := temp[n-1]
	return val
}

func getFinalState(nums []int, k int, multiplier int) []int {
	mod := 1000000007
	h := &MinHeap{}
	heap.Init(h)
	for i, val := range nums {
		heap.Push(h, item{val, i})
	}
	for i := 0; i < k; i++ {
		minItem := heap.Pop(h).(item)
		minVal := minItem.val % mod
		minVal *= multiplier
		heap.Push(h, item{minVal, minItem.index})
	}
	res := make([]int, len(nums))
	for h.Len() > 0 {
		minVal := heap.Pop(h).(item)
		res[minVal.index] = minVal.val % mod
	}
	return res
}

func main() {
	fmt.Printf("Res: %v", getFinalState([]int{100000, 2000}, 2, 1000000))
}
