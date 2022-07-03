package build_a_pile_of_cubes

import "math"

func FindNb(m int) int {
	total := 0
	iteration := 1

	for total <= m {
		total += int(math.Pow(float64(iteration), 3))

		if total == m {
			return iteration
		}

		iteration++
	}

	return -1
}
