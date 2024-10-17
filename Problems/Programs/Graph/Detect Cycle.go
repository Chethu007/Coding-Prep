package main

//
//import "fmt"
//
//// ListNode represents a node in the adjacency list
//type ListNode struct {
//	data int
//	next *ListNode
//}
//
//// Graph represents a graph using an adjacency list
//type Graph struct {
//	head []*ListNode
//}
//
//type item2 struct {
//	node   *ListNode
//	parent int
//}
//
//// createGraph initializes the graph with n nodes
//func createGraph(n int) *Graph {
//	g := &Graph{
//		head: make([]*ListNode, n),
//	}
//	return g
//}
//
//// addEdge adds an edge between two nodes u and v
//func (g *Graph) addEdge(u, v int) {
//	// Add edge from u to v
//	node := &ListNode{data: v, next: g.head[u]}
//	g.head[u] = node
//
//	// Add edge from v to u (since the graph is undirected)
//	node = &ListNode{data: u, next: g.head[v]}
//	g.head[v] = node
//}
//
//// isCycle detects if there is a cycle in the graph
//func isCycle(g *Graph, v, e int) bool {
//	visited := make([]int, v)
//	var q []item2
//	q = append(q, item2{g.head[0], -1})
//	for len(q) > 0 {
//		cur := q[0]
//		curNode := cur.node
//		q = q[1:]
//		for _, val := range g.head[curNode] {
//
//		}
//	}
//	return false
//}
//
//func main() {
//	n, m := 5, 5
//	graph := createGraph(n)
//
//	edges := [][]int{
//		{0, 1},
//		{1, 2},
//		{1, 4},
//		{2, 3},
//		{3, 4},
//	}
//
//	// Add edges to the graph
//	for _, edge := range edges {
//		graph.addEdge(edge[0], edge[1])
//	}
//
//	// Check if the graph contains a cycle
//	if isCycle(graph, n, m) {
//		fmt.Println("Graph contains a cycle")
//	} else {
//		fmt.Println("Graph does not contain a cycle")
//	}
//}
