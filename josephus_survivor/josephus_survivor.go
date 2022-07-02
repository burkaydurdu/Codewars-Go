package josephus_survivor

func JosephusSurvivor(n, k int) int {
	if n == 1 {
		return 1
	}
	return (JosephusSurvivor(n-1, k)+k-1)%n + 1
}
