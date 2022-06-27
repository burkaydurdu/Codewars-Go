package multiplication_table

func MultiplicationTable(size int) [][]int {
	var result [][]int

	for i := 1; i <= size; i++ {
		var row []int
		for j := i; j <= i*size; j = i + j {
			row = append(row, j)
		}
		result = append(result, row)
	}
	return result
}
