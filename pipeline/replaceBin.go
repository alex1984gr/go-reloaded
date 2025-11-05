package pipeline

import (
	"regexp"
	"strconv"
)

func ReplaceBin(input string) string {
	re := regexp.MustCompile(`\b[01]+\b`)
	return re.ReplaceAllStringFunc(input, func(bin string) string {
		if val, err := strconv.ParseInt(bin, 2, 64); err == nil {
			return strconv.FormatInt(val, 10)
		}
		return bin
	})
}
