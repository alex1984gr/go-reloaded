package pipeline

import (
	"fmt"
	"regexp"
	"strconv"
)

func ReplaceHex(lines []string) []string {
	re := regexp.MustCompile(`\b[0-9A-Fa-f]+\b`)

	var result []string

	for _, line := range lines {
		converted := re.ReplaceAllStringFunc(line, func(hexStr string) string {
			decValue, err := strconv.ParseInt(hexStr, 16, 64)
			if err != nil {
				return hexStr
			}
			return fmt.Sprintf("%d", decValue)
		})
		result = append(result, converted)
	}

	return result
}
