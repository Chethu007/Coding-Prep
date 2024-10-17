package main

import "fmt"

type MinStack struct {
	St    []int
	MinSt []int
}

func Constructor() MinStack {
	var st, minSt []int
	mS := MinStack{St: st, MinSt: minSt}
	return mS
}

func (this *MinStack) Push(val int) {
	this.St = append(this.St, val)
	n := len(this.MinSt)
	if n > 0 && val > this.MinSt[n-1] {
		this.MinSt = append(this.MinSt, this.MinSt[n-1])
	} else {
		this.MinSt = append(this.MinSt, val)
	}
}

func (this *MinStack) Pop() {
	if this.MinSt[len(this.MinSt)-1] == this.St[len(this.St)-1] {
		this.MinSt = this.MinSt[:len(this.MinSt)-1]
	}
	this.St = this.St[:len(this.St)-1]
}

func (this *MinStack) Top() int {
	return this.St[len(this.St)-1]
}

func (this *MinStack) GetMin() int {
	n := len(this.MinSt)
	if n == 0 {
		return -1
	}
	return this.MinSt[len(this.MinSt)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	obj := Constructor()
	fmt.Printf("GetMin: %v\n", obj.GetMin())
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Printf("GetMin: %v\n", obj.GetMin())
	obj.Pop()
	param_3 := obj.Top()
	fmt.Printf("Top: %v\n", param_3)
	param_4 := obj.GetMin()
	fmt.Printf("GetMin: %v\n", param_4)
}
