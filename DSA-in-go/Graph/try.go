package main


func solve(A int, B [][]int) int {
	adj := make([][]int, A)
	
	for i := range adj {
		adj[i] = make([]int, 0)
	}

	for i := range B {
		start := B[i][0]
		end := B[i][1]
		adj[start] = append(adj[start], end)
	}
	visited := make([]bool, A)

	if dfs(adj, 0, A-1, visited) {
		return 1
	}
	return 0
}

func dfs(adj [][]int, start, end int, visited []bool) bool {

	if start == end {
		return true
	}

	visited[start] = true

	for _, v := range adj[start] {
		if visited[v] == false {
			if dfs(adj, v, end, visited) {
				return true
			}
		}
	}

	return false
}
