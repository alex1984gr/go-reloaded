package pipeline

import (
	"strings"
)

// ApplyCaseTransform applies the selected case transformation ("upper", "lower", "capitalize", "none")
// to each string in the input slice and returns the transformed slice.
func ApplyCaseTranform(input []string, option string) []string {
	result := make([]string, len(input))
	for i, word := range input {
		switch strings.ToLower(option) {
		case "upper":
			result[i] = strings.ToUpper(word)
		case "lower":
			result[i] = strings.ToLower(word)
		case "capitalize": // capitalize first letter, keep the rest lowercase
			if len(word) > 0 {
				result[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
			} else {
				result[i] = word
			}
		case "none":
			result[i] = word
		default: // if the option is unknown, return the word unchanged
			result[i] = word
		}
	}
	return result
}
