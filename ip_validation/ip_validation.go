package ip_validation

import (
	"regexp"
	"strconv"
	"strings"
)

func Is_valid_ip(ip string) bool {
	str := strings.Split(ip, ".")

	for _, val := range str {
		if i, err := strconv.Atoi(val); err != nil || i > 255 {
			return false
		}
	}
	
	re, _ := regexp.Compile("([1-9]\\d{0,2}|0)\\.([1-9]\\d{0,2}|0)\\.([1-9]\\d{0,2}|0)\\.([1-9]\\d{0,2}|0)")
	matchIP := re.FindString(ip)

	return matchIP != ""
}
