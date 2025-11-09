package pipeline

import (
	"regexp"
	"strconv"
)

// ReplaceBin scans through a slice of strings and replaces any binary number with its decimal equivalent.
func ReplaceBin(input []string) []string {
	re := regexp.MustCompile(`\b[01]+\b`) // Create a regular expression that matches binary numbers (only 0s and 1s, as whole words)

	for i, word := range input { // Loop through every element in the input slice
		input[i] = re.ReplaceAllStringFunc(word, func(bin string) string { // Apply a function to every substring matching the regex
			if val, err := strconv.ParseInt(bin, 2, 64); err == nil { // Try to convert the binary string 'bin' to an integer (base 2)
				return strconv.FormatInt(val, 10) // If successful, convert that integer to a decimal string and return it
			}
			return bin // If there’s an error, just return the original string unchanged
		})
	}
	return input // Return the modified slice with all binary numbers replaced
}
