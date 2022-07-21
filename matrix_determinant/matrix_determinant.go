package matrix_determinant

func subArray(matrix [][]int, index int) [][]int {
	var subMatrix = make([][]int, len(matrix)-1)

	subIndex := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if j != index && i != 0 {
				subMatrix[subIndex] = append(subMatrix[subIndex], matrix[i][j])
			}
		}

		if i != 0 {
			subIndex++
		}
	}

	return subMatrix
}

func Determinant(matrix [][]int) int {
	if len(matrix) == 1 {
		return matrix[0][0]
	}

	if len(matrix) == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	}

	total := 0
	factor := 1
	for i := 0; i < len(matrix); i++ {
		if i%2 != 0 {
			factor = -1
		} else {
			factor = 1
		}

		total += factor * matrix[0][i] * Determinant(subArray(matrix, i))
	}

	return total
}
