package pipeline

import (
	"fmt"
	"regexp"
	"strconv"
)

// ReplaceHex replaces hexadecimal numbers (like "1E (hex)") with their decimal equivalents
func ReplaceHex(lines []string) []string {
	re := regexp.MustCompile(`\b([0-9A-Fa-f]+)\s*\(hex\)`) // Matches tokens like "1E (hex)" or "FF(hex)"

	var result []string // This will hold the processed output lines

	for _, line := range lines { // Iterate through each line in the input
		converted := re.ReplaceAllStringFunc(line, func(match string) string { // For each hex match found, apply a conversion
			hexMatch := regexp.MustCompile(`^[0-9A-Fa-f]+`).FindString(match) // Extract only the hexadecimal part before "(hex)"
			if hexMatch == "" {                                               // If no valid hex number found
				return match // Return the original text unchanged
			}

			decValue, err := strconv.ParseInt(hexMatch, 16, 64) // Convert from base 16 (hex) to base 10 (decimal)
			if err != nil {                                     // If conversion fails
				return match // Keep the original match
			}

			return fmt.Sprintf("%d", decValue) // Return the converted decimal number as a string
		})
		result = append(result, converted) // Add the converted line to the final result
	}

	return result // Return the list of processed lines
}
