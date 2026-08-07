func spiralOrder(matrix [][]int) []int {
    m := len(matrix)
	n := len(matrix[0])
	result := make([]int, 0, m*n)
	var x, y int
	var mx, my int = 1, 0
	var minX, maxX int = -1, n
	var minY, maxY int = -1, m
	for i := 0; i < m*n; i++ {
		result = append(result, matrix[y][x])
		ty := y + my
		tx := x + mx
		if tx == minX || tx == maxX || ty == minY || ty == maxY {
			mx, my = -my, mx
			if tx == maxX {
				minY++
			} else if ty == maxY {
				maxX--
			} else if tx == minX {
				maxY--
			} else {
				minX++
			}
		}
		x += mx
		y += my
	}
	return result
}
