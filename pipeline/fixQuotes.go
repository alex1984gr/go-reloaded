package pipeline

import (
	"strings"
)

// FixQuotes handles single quotes around words or phrases
// Example:
//
//	Input:  [I am ' awesome ']
//	Output: [I am 'awesome']
func FixQuotes(tokens []string) []string {
	openIndex := -1

	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "'" {
			if openIndex == -1 {
				// opening quote found
				openIndex = i
			} else {
				// closing quote found
				// collect and trim inner tokens
				var inner []string
				for j := openIndex + 1; j < i; j++ {
					inner = append(inner, strings.TrimSpace(tokens[j]))
				}
				// join inner content with single spaces
				content := strings.Join(inner, " ")
				// build quoted token
				quoted := "'" + content + "'"

				// rebuild tokens: tokens before openIndex + quoted + tokens after i
				newTokens := make([]string, 0, len(tokens)-(i-openIndex))
				newTokens = append(newTokens, tokens[:openIndex]...)
				newTokens = append(newTokens, quoted)
				if i+1 < len(tokens) {
					newTokens = append(newTokens, tokens[i+1:]...)
				}

				tokens = newTokens

				// reset scanning index and openIndex
				i = openIndex
				openIndex = -1
			}
		}
	}

	return tokens
}
