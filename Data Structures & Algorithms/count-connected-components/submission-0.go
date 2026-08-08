func countComponents(n int, edges [][]int) int {
    graph := make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	visited := make([]bool, n)
	var dfs func(v int)
	dfs = func(v int) {
		for _, next := range graph[v] {
			if !visited[next] {
				visited[next] = true
				dfs(next)
			}
		}
	}
	count := 0
	for i := range visited {
		if !visited[i] {
			count++
			visited[i] = true
			dfs(i)
		}
	}
	return count
}
