package triangle

import "math"

func IsTriangle(a, b, c int) bool {
	return (math.Abs(float64(b-c)) < float64(a)) && a < b+c
}
