package sum_by_factors

import "fmt"

func getMax(vals []int) int {
	max := 0
	for _, v := range vals {
		if v > max {
			max = v
		}

		if v*-1 > max {
			max = v * -1
		}
	}

	return max
}

func isPrime(n, i int) bool {

	if n < 2 {
		return false
	}

	if n == 2 {
		return true
	}

	if n%i == 0 {
		return false
	}
	if i*i > n {
		return true
	}

	// Check for next divisor
	return isPrime(n, i+1)
}

func getPrimes(limit int) []int {
	var primes []int
	for i := 2; i <= limit; i++ {
		if isPrime(i, 2) {
			primes = append(primes, i)
		}
	}

	return primes
}

func SumOfDivided(lst []int) string {
	var data = make(map[int]int)
	max := getMax(lst)
	primes := getPrimes(max)

	for _, prime := range primes {
		for _, l := range lst {
			if l%prime == 0 {
				data[prime] += l
			}
		}
	}

	var result string

	//for k, v := range data {
	//	result += fmt.Sprintf("(%d %d)", k, v)
	//}

	for _, prime := range primes {
		if d, ok := data[prime]; ok {
			result += fmt.Sprintf("(%d %d)", prime, d)
		}
	}

	return result
}
