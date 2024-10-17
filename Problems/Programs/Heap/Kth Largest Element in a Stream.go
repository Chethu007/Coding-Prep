package main

import "sort"

type KthLargest struct {
	k   int
	arr []int
}

func Constructor(k int, nums []int) KthLargest {
	kL := KthLargest{k: k, arr: nums}
	return kL
}

func (this *KthLargest) Add(val int) int {
	this.arr = append(this.arr, val)
	sort.Ints(this.arr)
	return this.arr[len(this.arr)-this.k]
}

func main() {
	nums := []int{4, 5, 8, 2}
	obj := Constructor(3, nums)
	param_1 := obj.Add(3)
	println("param_1: ", param_1)
	param_1 = obj.Add(5)
	println("param_2: ", param_1)
	param_1 = obj.Add(10)
	println("param_3: ", param_1)
	param_1 = obj.Add(9)
	println("param_4: ", param_1)
	param_1 = obj.Add(4)
	println("param_5: ", param_1)
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
