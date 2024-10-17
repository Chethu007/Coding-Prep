package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	m := make(map[int]int)
	for _, val := range nums {
		m[val] = 1
	}
	for head != nil && m[head.Val] == 1 {
		head = head.Next
	}
	cur := head
	for cur != nil && cur.Next != nil {
		if m[cur.Next.Val] == 1 {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func deleteAll(num int, head *ListNode) *ListNode {
	for head != nil && head.Val == num {
		head = head.Next
	}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == num {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
