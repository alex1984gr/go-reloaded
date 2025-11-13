package pipeline // This file belongs to the "pipeline" package

import (
	"fmt"     // Used for debug printing if needed
	"strconv" // Needed to parse binary strings to integers
	"strings" // Needed for string comparison
)

// ReplaceBin scans tokens for "(bin)" markers and converts the previous token
// from binary (base 2) to decimal (base 10). Invalid binaries are ignored.
func ReplaceBin(tokens []string) []string {
	var result []string // Holds the processed tokens

	for i := 0; i < len(tokens); i++ { // Iterate through each token
		token := tokens[i]

		// Check if the current token is "(bin)" (case-insensitive)
		if strings.EqualFold(token, "(bin)") {
			if len(result) > 0 { // Make sure there is a previous token to convert
				binWord := result[len(result)-1] // Get the previous token

				// Try converting from binary string to decimal integer
				value, err := strconv.ParseInt(binWord, 2, 64)
				if err == nil { // Conversion succeeded
					result[len(result)-1] = fmt.Sprintf("%d", value) // Replace with decimal
				}
				// If conversion fails, leave the word unchanged
			}
			// Skip the "(bin)" token itself
			continue
		}

		// For normal words, just append to the result
		result = append(result, token)
	}

	return result // Return the transformed tokens
}
