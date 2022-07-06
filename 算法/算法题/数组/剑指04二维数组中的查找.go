package aboutArray

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 {
		return false
	}
	col := len(matrix[0]) - 1
	row := 0
	for row < len(matrix) && col >= 0 {
		if target > matrix[row][col] {
			row++
		} else if target < matrix[row][col] {
			col--
		} else {
			return true
		}
	}
	return false
}
