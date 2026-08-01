func exist(board [][]byte, word string) bool {
    m := len(board)
    n := len(board[0])
    visited := make([][]bool, m)
    for i := range visited {
        visited[i] = make([]bool, n)
    }
    dx := []int{0, 1, 0, -1}
    dy := []int{1, 0, -1, 0}
    var backtracking func(x, y, i int) bool 
    backtracking = func(x, y, i int) bool {
        if i == len(word) {
            return true
        }
        for j := range dx {
            tx := x + dx[j]
            ty := y + dy[j]
            if ty >= 0 && ty < m && tx >= 0 && tx < n && !visited[ty][tx] && board[ty][tx] == word[i] {
                visited[ty][tx] = true
                if backtracking(tx, ty, i+1) {
                    return true
                }
                visited[ty][tx] = false
            }
        }
        return false
    }
    for y := 0; y < m; y++ {
        for x := 0; x < n; x++ {
            if board[y][x] != word[0] {
                continue
            }
            visited[y][x] = true
            if backtracking(x, y, 1) {
                return true
            }
            visited[y][x] = false
        }
    }
    return false 
}
