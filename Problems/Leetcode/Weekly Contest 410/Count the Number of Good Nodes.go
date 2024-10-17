package main

import "fmt"

func countGoodNodes(edges [][]int) int {
	n := len(edges)
	adj := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		adj[i] = make([]int, n+1)
	}
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			adj[i][j] = -1
		}
	}
	for _, val := range edges {
		adj[val[0]] = append(adj[val[0]], val[1])
		adj[val[1]] = append(adj[val[1]], val[0])
	}
	//fmt.Printf("Adj: %v\n", adj)
	visited := make([]int, n+1)
	temp := make([]int, n+1)
	for i := n; i >= 0; i-- {
		count := 0
		for _, val := range adj[i] {
			if val == -1 {
				continue
			}
			if visited[val] == 0 {
				count++
			} else {
				count += temp[val]
			}
		}
		visited[i] = 1
		temp[i] = count
	}
	fmt.Printf("Arr: %v\n", temp)
	res := 0
	visited1 := make([]int, n+1)
	for i := 0; i <= n; i++ {
		var temp1 []int
		flag := 1
		for _, val := range adj[i] {
			if val == -1 {
				continue
			}
			if visited1[val] == 0 {
				temp1 = append(temp1, temp[val])
			}
		}
		visited1[i] = 1
		if len(temp1) <= 1 {
			res++
			continue
		}
		for j := 1; j < len(temp1); j++ {
			if temp1[j] != temp1[j-1] {
				flag = 0
			}
		}
		if flag == 1 {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(countGoodNodes([][]int{{4, 0}, {6, 4}, {3, 6}, {5, 0}, {1, 5}, {2, 1}}))
}
