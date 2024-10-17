package Binary_Search_Tree

//
//import "fmt"
//
//type Node struct {
//	data  int
//	left  *Node
//	right *Node
//}
//
//func NewNode(data int) *Node {
//	return &Node{data: data, left: nil, right: nil}
//}
//
//func (n *Node) Insert(data int) {
//	if data <= n.data {
//		if n.left == nil {
//			n.left = NewNode(data)
//		} else {
//			n.left.Insert(data)
//		}
//	} else {
//		if n.right == nil {
//			n.right = NewNode(data)
//		} else {
//			n.right.Insert(data)
//		}
//	}
//}
//
//// Search searches for a value in the binary tree
//func (n *Node) Search(data int) bool {
//	if n == nil {
//		return false
//	}
//	if n.data == data {
//		return true
//	} else if data < n.data {
//		return n.left.Search(data)
//	} else {
//		return n.right.Search(data)
//	}
//}
//
//func (n *Node) InOrder() {
//	if n != nil {
//		n.left.InOrder()
//		fmt.Print(n.data, " ")
//		n.right.InOrder()
//	}
//}
//
//func main() {
//	// Creating the root of the binary tree
//	root := NewNode(10)
//
//	// Inserting values
//	values := []int{5, 3, 7, 15, 12, 17}
//	for _, value := range values {
//		root.Insert(value)
//	}
//
//	//// Searching for values
//	fmt.Println("Search 7:", root.Search(17)) // true
//	//fmt.Println("Search 100:", root.Search(100)) // false
//	//
//	//// Traversing the tree
//	fmt.Print("InOrder: ")
//	root.InOrder()
//	fmt.Println()
//	//
//	//fmt.Print("PreOrder: ")
//	//root.PreOrder()
//	//fmt.Println()
//	//
//	//fmt.Print("PostOrder: ")
//	//root.PostOrder()
//	//fmt.Println()
//}
