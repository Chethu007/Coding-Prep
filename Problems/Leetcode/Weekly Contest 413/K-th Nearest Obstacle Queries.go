package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func absVal(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func resultsArray(queries [][]int, k int) []int {
	h := &IntHeap{}
	heap.Init(h)
	res := make([]int, len(queries))
	for i, q := range queries {
		if i+1 < k {
			heap.Push(h, absVal(q[0])+absVal(q[1]))
			res[i] = -1
		} else if i+1 == k {
			heap.Push(h, absVal(q[0])+absVal(q[1]))
			res[i] = (*h)[0]
		} else {
			top, val := (*h)[0], absVal(q[0])+absVal(q[1])
			if top > val {
				heap.Pop(h)
				heap.Push(h, absVal(val))
				res[i] = (*h)[0]
			} else {
				res[i] = (*h)[0]
			}
		}
	}
	return res
}

func main() {
	queries := [][]int{{6, 10}, {0, -10}, {2, -6}}

	fmt.Printf("Res: %v\n", resultsArray(queries, 2))
}
