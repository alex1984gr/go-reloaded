package pipeline

import (
	"regexp"
	"strconv"
)

func ReplaceBin(input []string) []string {
	re := regexp.MustCompile(`\b[01]+\b`)

	for i, word := range input {
		input[i] = re.ReplaceAllStringFunc(word, func(bin string) string {
			if val, err := strconv.ParseInt(bin, 2, 64); err == nil {
				return strconv.FormatInt(val, 10)
			}
			return bin
		})
	}
	return input
}
