package count_ip_addresses

import (
	"strconv"
	"strings"
)

func stringArrToIntegerArr(sArr []string) []int {
	var iArr = make([]int, len(sArr))

	for idx, i := range sArr {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		iArr[idx] = j
	}

	return iArr
}

func IpsBetween(start, end string) int {
	firstIPSteps := stringArrToIntegerArr(strings.Split(start, "."))
	lastIPSteps := stringArrToIntegerArr(strings.Split(end, "."))

	return ((lastIPSteps[0] - firstIPSteps[0]) * 256 * 256 * 256) + ((lastIPSteps[1] - firstIPSteps[1]) * 256 * 256) + ((lastIPSteps[2] - firstIPSteps[2]) * 256) + (lastIPSteps[3] - firstIPSteps[3])
}
