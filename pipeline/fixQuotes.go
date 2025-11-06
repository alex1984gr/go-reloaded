package pipeline

import "strings"

func FixQuotes(input []string) []string {
	result := []string{}
	openQuote := false
	for _, token := range input {
		trimmed := strings.TrimSpace(token)
		if strings.Contains(trimmed, "\"") || strings.Contains(trimmed, "'") {
			if openQuote {
				if len(result) > 0 {
					last := result[len(result)-1]
					result[len(result)-1] = strings.TrimSpace(last + " " + trimmed)
				} else {
					result = append(result, trimmed)
				}
				openQuote = false
			} else {
				result = append(result, trimmed)
				openQuote = true
			}
		} else {
			result = append(result, trimmed)
		}
	}
	return result
}
