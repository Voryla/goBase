package aboutArray

func spiralOrder(matrix [][]int) []int {
	answer := make([]int, 0)
	rowMax := len(matrix) - 1
	colMax := len(matrix[0]) - 1
	lowRow, lowCol := 0, 0
	for {
		for i := lowCol; i <= colMax; i++ {
			answer = append(answer, matrix[lowRow][i]) //从左往右
		}
		if lowRow++; lowRow > rowMax {
			break
		}
		for i := lowRow; i <= rowMax; i++ {
			answer = append(answer, matrix[i][colMax]) //从上往下
		}
		if colMax--; colMax < lowCol {
			break
		}
		for i := colMax; i >= lowCol; i-- {
			answer = append(answer, matrix[rowMax][i]) //从右往左
		}
		if rowMax--; rowMax < lowRow {
			break
		}
		for i := rowMax; i >= lowRow; i-- {
			answer = append(answer, matrix[i][lowCol]) //从下往上
		}
		if lowCol++; lowCol > colMax {
			break
		}
	}
	return answer
}
