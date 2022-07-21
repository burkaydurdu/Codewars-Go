package sum_of_intervals

func remove(slice [][2]int, s int) [][2]int {
	return append(slice[:s], slice[s+1:]...)
}

func SumOfIntervals(intervals [][2]int) int {
	pointer := 0
	isReturn := false

	for pointer < len(intervals) {
		for j := pointer + 1; j < len(intervals); j++ {
			if intervals[j][0] < intervals[pointer][0] && intervals[j][1] > intervals[pointer][0] && intervals[j][1] < intervals[pointer][1] {
				intervals[pointer] = [2]int{intervals[j][0], intervals[pointer][1]}
				intervals = remove(intervals, j)
				isReturn = true
				break
				// left side
			} else if intervals[pointer][0] < intervals[j][0] && intervals[pointer][1] > intervals[j][0] && intervals[pointer][1] < intervals[j][1] {
				intervals[pointer] = [2]int{intervals[pointer][0], intervals[j][1]}
				intervals = remove(intervals, j)
				isReturn = true
				break
			} else if intervals[pointer][0] >= intervals[j][0] && intervals[pointer][1] <= intervals[j][1] {
				intervals = remove(intervals, pointer)
				isReturn = true
				break
				// inside
			} else if intervals[j][0] >= intervals[pointer][0] && intervals[j][1] <= intervals[pointer][1] {
				intervals = remove(intervals, j)
				isReturn = true
				break
			}
		}

		if isReturn {
			pointer = 0
			isReturn = false
		} else {
			pointer++
		}
	}

	total := 0
	for _, i := range intervals {
		total += i[1] - i[0]
	}

	return total
}
