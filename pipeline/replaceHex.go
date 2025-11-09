package pipeline

import (
	"fmt"
	"regexp"
	"strconv"
)

func ReplaceHex(lines []string) []string {
	re := regexp.MustCompile(`\b([0-9A-Fa-f]+)\s*\(hex\)`)

	var result []string

	for _, line := range lines {
		converted := re.ReplaceAllStringFunc(line, func(match string) string {
			hexMatch := regexp.MustCompile(`^[0-9A-Fa-f]+`).FindString(match)
			if hexMatch == "" {
				return match
			}

			decValue, err := strconv.ParseInt(hexMatch, 16, 64)
			if err != nil {
				return match
			}

			return fmt.Sprintf("%d", decValue)
		})
		result = append(result, converted)
	}

	return result
}
