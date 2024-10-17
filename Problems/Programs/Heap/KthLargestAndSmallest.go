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

//func kthLargest(nums []int, k int) int {
//	h := &IntHeap{}
//	heap.Init(h)
//	for i := range nums {
//		heap.Push(h, nums[i])
//		if h.Len() > k {
//			heap.Pop(h)
//		}
//	}
//	//fmt.Printf("%dth Largest is : %d\n", k, heap.Pop(h))
//	return heap.Pop(h).(int)
//}

func kthSmallest(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)
	for i := range nums {
		heap.Push(h, nums[i])
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	//fmt.Printf("%dth Largest is : %d\n", k, heap.Pop(h))
	return heap.Pop(h).(int)
}

func kthSmallestElements(nums []int, k int) []int {
	h := &IntHeap{}
	heap.Init(h)
	for i := range nums {
		heap.Push(h, nums[i])
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	var res []int
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(int))
	}
	return res
}

func main() {
	nums := []int{2, 3, 4, 3, 6, 7, 8}
	k := 3
	fmt.Println(nums)
	//largest := kthLargest(nums, k)
	//fmt.Printf("%dth Largest is : %d\n", k, largest)
	smallest := kthSmallest(nums, k)
	fmt.Printf("%dth Smallest is : %d\n", k, smallest)
	smallestEle := kthSmallestElements(nums, k)
	fmt.Printf("%d Smallest Elements is : %v\n", k, smallestEle)
}
