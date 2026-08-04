func validTree(n int, edges [][]int) bool {
	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	visited := make([]bool, n)
	var dfs func(node, from int) bool
	dfs = func(node, from int) bool {
		for _, next := range graph[node] {
			if next == from {
				continue
			}
			if visited[next] {
				return false
			}
			visited[next] = true
			if !dfs(next, node) {
				return false
			}
		}
		return true
	}
	visited[0] = true
	if !dfs(0, -1) {
		return false
	}
	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}
