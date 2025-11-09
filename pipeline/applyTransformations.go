package pipeline

import (
	"fmt"
)

// ApplyTransformations εφαρμόζει όλη την αλυσίδα μετασχηματισμών
func ApplyTransformations(tokens []string) []string {
	fmt.Printf("DEBUG: Original tokens: %v\n", tokens)

	tokens = FixArticles(tokens)
	fmt.Printf("DEBUG: After FixArticles: %v\n", tokens)

	tokens = FixQuotes(tokens)
	fmt.Printf("DEBUG: After FixQuotes: %v\n", tokens)

	tokens = ReplaceBin(tokens)
	fmt.Printf("DEBUG: After ReplaceBin: %v\n", tokens)

	tokens = ReplaceHex(tokens)
	fmt.Printf("DEBUG: After ReplaceHex: %v\n", tokens)

	tokens = ApplyCaseTransform(tokens)
	fmt.Printf("DEBUG: After ApplyCaseTransform: %v\n", tokens)

	tokens = FormatPunctuation(tokens)
	fmt.Printf("DEBUG: After FormatPunctuation: %v\n", tokens)

	return tokens
}
