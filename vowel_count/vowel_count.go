package vowel_count

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func GetCount(str string) (count int) {
	vowel := []string{"a", "e", "i", "o", "u"}

	for _, i := range str {
		if contains(vowel, string(i)) {
			count++
		}
	}

	return count
}
