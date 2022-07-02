package rgb_to_hex_conversion

import (
	"fmt"
	"strings"
)

func RGB(r, g, b int) string {
	var rHex, gHex, bHex = "00", "00", "00"

	if r > 255 {
		rHex = "FF"
	} else if r > 0 {
		rHex = fmt.Sprintf("%02x", r)
	}

	if g > 255 {
		gHex = "FF"
	} else if g > 0 {
		gHex = fmt.Sprintf("%02x", g)
	}

	if b > 255 {
		bHex = "FF"
	} else if b > 0 {
		bHex = fmt.Sprintf("%02x", b)
	}

	return strings.ToUpper(rHex + gHex + bHex)
}
