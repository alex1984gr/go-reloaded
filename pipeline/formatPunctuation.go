package pipeline

import "strings"

func FormatPunctuation(tokens []string) []string {
	if len(tokens) == 0 {
		return tokens
	}
	result := []string{}
	for i, token := range tokens {
		if isPunctuation(token) && len(result) > 0 {
			result[len(result)-1] += token
		} else if i > 0 && isPunctuation(tokens[i-1]) {
			result = append(result, token)
		} else {
			result = append(result, token)
		}
	}
	output := strings.Join(result, " ")
	return strings.Fields(output)
}
func isPunctuation(token string) bool {
	switch token {
	case ".", ",", "!", "?", ";", ":":
		return true
	default:
		return false
	}
}
