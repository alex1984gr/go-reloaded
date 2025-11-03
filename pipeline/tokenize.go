package pipeline

import "strings"

// Tokenize χωρίζει το κείμενο σε λέξεις με βάση τα κενά.
func Tokenize(input []string) []string {
	var result []string
	for _, line := range input {
		words := strings.Fields(line)
		result = append(result, words...)
	}
	return result
}
