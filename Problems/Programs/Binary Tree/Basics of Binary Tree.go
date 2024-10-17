package main

import (
	"container/list"
	"fmt"
)

// Node represents a node in the binary tree
type Node struct {
	data  int
	left  *Node
	right *Node
}

// NewNode creates a new node
func NewNode(data int) *Node {
	return &Node{data: data, left: nil, right: nil}
}

// Insert inserts a new node into the binary tree in level order
func (n *Node) Insert(data int) {
	newNode := NewNode(data)
	if n == nil {
		*n = *newNode
		return
	}

	// Create a queue and enqueue the root node
	q := list.New()
	q.PushBack(n)

	// Do level order traversal until we find an empty place
	for q.Len() > 0 {
		// Get the front node of the queue
		node := q.Remove(q.Front()).(*Node)

		// Check left child
		if node.left == nil {
			node.left = newNode
			return
		} else {
			q.PushBack(node.left)
		}

		// Check right child
		if node.right == nil {
			node.right = newNode
			return
		} else {
			q.PushBack(node.right)
		}
	}
}

// InOrder traverses the tree in in-order fashion
func (n *Node) InOrder() {
	if n != nil {
		n.left.InOrder()
		fmt.Print(n.data, " ")
		n.right.InOrder()
	}
}

// PreOrder traverses the tree in pre-order fashion
func (n *Node) PreOrder() {
	if n != nil {
		fmt.Print(n.data, " ")
		n.left.PreOrder()
		n.right.PreOrder()
	}
}

// PostOrder traverses the tree in post-order fashion
func (n *Node) PostOrder() {
	if n != nil {
		n.left.PostOrder()
		n.right.PostOrder()
		fmt.Print(n.data, " ")
	}
}

func levelOrderTraversal(root *Node) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	q := list.New()
	q.PushBack(root)
	for q.Len() > 0 {
		n := q.Len()
		var temp []int
		for i := 0; i < n; i++ {
			var node *Node
			//node = q.Remove(q.Front()).(*Node)

			node = q.Front().Value.(*Node)
			q.Remove(q.Front())
			temp = append(temp, node.data)
			if node.left != nil {
				q.PushBack(node.left)
			}
			if node.right != nil {
				q.PushBack(node.right)
			}
		}
		res = append(res, temp)
	}
	return res
}

func iterativePreOrderTraversal(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}
	var st []*Node
	st = append(st, root)
	for len(st) > 0 {
		node := st[len(st)-1]
		st = st[:len(st)-1]
		res = append(res, node.data)
		if node.right != nil {
			st = append(st, node.right)
		}
		if node.left != nil {
			st = append(st, node.left)
		}
	}
	return res
}

//	func iterativeInOrderTraversal(root *Node) []int {
//		var res []int
//		var st []*Node
//		node := root
//		for {
//			if node != nil {
//				st = append(st, node)
//				node = node.left
//			} else {
//				if len(st) == 0 {
//					break
//				}
//				node = st[len(st)-1]
//				st = st[:len(st)-1]
//				res = append(res, node.data)
//				node = node.right
//			}
//		}
//		return res
//	}
func iterativeInOrderTraversal(root *Node) []int {
	var res []int
	var st []*Node
	node := root
	for node != nil || len(st) > 0 {
		for node != nil {
			st = append(st, node)
			node = node.left
		}
		node = st[len(st)-1]
		st = st[:len(st)-1]
		res = append(res, node.data)
		node = node.right
	}
	return res
}

func main() {
	// Creating the root of the binary tree
	root := NewNode(10)

	// Inserting values
	//values := []int{5, 3, 7, 15, 12, 17}
	//for _, value := range values {
	//	root.Insert(value)
	//}
	root.left = NewNode(11)
	root.right = NewNode(12)
	root.left.left = NewNode(13)
	root.left.right = NewNode(14)
	root.right.left = NewNode(15)
	root.right.right = NewNode(16)

	fmt.Printf("levelOrderTraversal: %v\n", levelOrderTraversal(root))
	// Traversing the tree
	fmt.Print("InOrder: ")
	root.InOrder()
	fmt.Println()

	fmt.Printf("iterative InOrder Traversal: %v\n", iterativeInOrderTraversal(root))

	fmt.Print("PreOrder: ")
	root.PreOrder()
	fmt.Println()

	fmt.Printf("iterative PreOrder Traversal: %v\n", iterativePreOrderTraversal(root))

	fmt.Print("PostOrder: ")
	root.PostOrder()
	fmt.Println()
}
