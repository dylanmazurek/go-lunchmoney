package truncate

import (
	"unicode"
)

func TruncateText(s string, max int) string {
	lastSpaceIdx := -1
	len := 0
	for i, r := range s {
		if unicode.IsSpace(r) {
			lastSpaceIdx = i
		}

		len++
		if len >= max {
			if lastSpaceIdx != -1 {
				return s[:lastSpaceIdx] + "..."
			}
		}
	}

	return s
}
