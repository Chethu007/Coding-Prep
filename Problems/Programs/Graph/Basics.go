package main

import "fmt"

func adjacencyList(matrix [][]int, n int) [][]int {
	adj := make([][]int, n+1)
	for i := 0; i < len(matrix); i++ {
		adj[matrix[i][0]] = append(adj[matrix[i][0]], matrix[i][1])
		adj[matrix[i][1]] = append(adj[matrix[i][1]], matrix[i][0])
	}
	return adj
}
func adjacencyMatrix(matrix [][]int, n int) [][]int {
	adj := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		adj[i] = make([]int, n+1)
	}
	for i := 0; i < len(matrix); i++ {
		adj[matrix[i][0]][matrix[i][1]] = 1
		adj[matrix[i][1]][matrix[i][0]] = 1
	}
	return adj
}

func weightedAdjacencyList(matrix [][]int, n int) [][]int {
	adj := make([][]int, n+1)
	for i := 0; i < len(matrix); i++ {
		adj[matrix[i][0]] = append(adj[matrix[i][0]], matrix[i][1])
		adj[matrix[i][1]] = append(adj[matrix[i][1]], matrix[i][0])
	}
	return adj
}
func weightedAdjacencyMatrix(matrix [][]int, n int) [][]int {
	adj := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		adj[i] = make([]int, n+1)
	}
	for i := 0; i < len(matrix); i++ {
		adj[matrix[i][0]][matrix[i][1]] = 1
		adj[matrix[i][1]][matrix[i][0]] = 1
	}
	return adj
}

func BFS(v int, adj [][]int) []int {
	var res []int
	q := []int{1}
	visited := make([]int, v+1)
	visited[q[0]] = 1
	for len(q) > 0 {
		temp := q[0]
		q = q[1:]
		for _, val := range adj[temp] {
			if visited[val] != 1 {
				q = append(q, val)
				visited[val] = 1
			}
		}
		res = append(res, temp)
	}
	return res
}

func DFS(v int, adj [][]int, res *[]int, visited []int) {
	*res = append(*res, v)
	visited[v] = 1
	for _, val := range adj[v] {
		if visited[val] != 1 {
			DFS(val, adj, res, visited)
			visited[val] = 1
		}
	}
}

func main() {
	//matrix := [][]int{{1, 2}, {1, 3}, {3, 4}, {2, 4}, {2, 5}, {4, 5}}
	//adj := adjacencyList(matrix, 5)
	//fmt.Printf("Adjacency List: %v\n", adj)
	//adj = adjacencyMatrix(matrix, 5)
	//fmt.Printf("Adjacency Matrix: %v\n", adj)
	//adj := [][]int{{1, 2, 3}, {}, {4}, {}, {}}
	//fmt.Printf("Res: %v", BFS(5, adj))
	//adj := [][]int{{}, {2, 6}, {1, 3, 4}, {2}, {2, 5}, {4, 8}, {1, 7, 9}, {6, 8}, {5, 7}, {6}}
	//fmt.Printf("Res: %v", BFS(9, adj))
	adj := [][]int{{}, {2, 3}, {1, 5, 6}, {1, 4, 7}, {3, 8}, {2}, {2}, {3, 8}, {4, 7}}
	var res []int
	visited := make([]int, 9)
	DFS(3, adj, &res, visited)
	fmt.Printf("Res: %v", res)
}
