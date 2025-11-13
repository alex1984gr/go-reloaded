package pipeline

import (
	"fmt"
	"strconv"
	"strings"
)

// ReplaceHex scans through all tokens and whenever it finds "(hex)",
// it replaces the *previous word* (which is always a hexadecimal number)
// with its decimal equivalent.
func ReplaceHex(tokens []string) []string {
	var result []string // Holds the final list of processed tokens

	for i := 0; i < len(tokens); i++ { // Iterate through each token
		token := tokens[i]

		// Check if the current token is "(hex)"
		if strings.EqualFold(token, "(hex)") {
			// Make sure there's a previous word to convert
			if len(result) > 0 {
				// Take the last word added to the result slice
				hexWord := result[len(result)-1]

				// Try converting it from hexadecimal (base 16) to decimal (base 10)
				value, err := strconv.ParseInt(hexWord, 16, 64)
				if err == nil {
					// If conversion succeeded, replace the previous word with the decimal number
					result[len(result)-1] = fmt.Sprintf("%d", value)
				}
			}
			// Skip adding "(hex)" itself to the result
			continue
		}

		// If it's a normal word, just add it to the result
		result = append(result, token)
	}

	// Return the transformed list of tokens
	return result
}
