package how_many_numbers

import (
	"strconv"
	"strings"
)

func splitInt(n int) []int {
	var slc []int
	for n > 0 {
		slc = append(slc, n%10)
		n = n / 10
	}
	return slc
}

func FindAll(sumDig, digs int) []int {
	minValue, _ := strconv.Atoi(strings.Repeat("1", digs))
	maxValue, _ := strconv.Atoi(strings.Repeat("9", digs))
	var validArr []int

	for i := minValue; i <= maxValue; i++ {
		sliceVal := splitInt(i)
		min := 10
		total := 0
		invalid := false

		for _, v := range sliceVal {
			if v > min {
				invalid = true
				break
			}

			min = v
			total += v
		}

		if !invalid && total == sumDig {
			validArr = append(validArr, i)
		}
	}

	validArrCount := len(validArr)
	if validArrCount > 0 {
		return []int{validArrCount, validArr[0], validArr[validArrCount-1]}
	}

	return []int{}
}
