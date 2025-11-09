package pipeline

import (
	"fmt"
)

// ApplyTransformations applies all transformations to a slice of tokens
func ApplyTransformations(tokens []string) []string {
	if len(tokens) == 0 {
		return tokens
	}

	fmt.Println("DEBUG: Original tokens:", tokens)
	tokens = FixArticles(tokens)
	fmt.Println("DEBUG: After FixArticles:", tokens)

	tokens = FixQuotes(tokens)
	fmt.Println("DEBUG: After FixQuotes:", tokens)

	tokens = ReplaceBin(tokens)
	fmt.Println("DEBUG: After ReplaceBin:", tokens)

	tokens = ReplaceHex(tokens)
	fmt.Println("DEBUG: After ReplaceHex:", tokens)

	tokens = ApplyCaseTransform(tokens)
	fmt.Println("DEBUG: After ApplyCaseTransform:", tokens)

	tokens = FormatPunctuation(tokens)
	fmt.Println("DEBUG: After FormatPunctuation:", tokens)

	return tokens
}
