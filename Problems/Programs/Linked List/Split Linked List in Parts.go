package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	count := 0
	node := head
	for node != nil {
		node = node.Next
		count++
	}
	totalSubNode, extraNode := count/k, count%k
	res := make([]*ListNode, k)
	cur := head
	idx := 0
	for cur != nil {
		curLength := totalSubNode
		if extraNode > 0 {
			curLength++
			extraNode--
		}
		res[idx] = cur
		for i := 0; i < curLength-1; i++ {
			cur = cur.Next
		}
		temp := cur.Next
		cur.Next = nil
		cur = temp
		idx++
	}
	return res
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	println()
}

func main() {
	//newList := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, &ListNode{8, &ListNode{9, &ListNode{10, nil}}}}}}}}}}
	//newList := nil
	res := splitListToParts(nil, 5)
	for _, node := range res {
		if node == nil {
			println("Empty")
		} else {
			printList(node)
		}
	}
	//fmt.Printf("Res: %v\n", )
}
