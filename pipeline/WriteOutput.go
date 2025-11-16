package pipeline

import (
	"os"
	"strings"
)

func JoinTokens(tokens []string) string {
	// Join tokens with a single space. Punctuation tokens are attached to
	// words by FormatPunctuation, so joining with spaces yields properly
	// spaced output (e.g. "Hello, world!").
	return strings.Join(tokens, " ")
}

func WriteOutput(filename string, input []string) error {
	output := strings.Join(input, " ")
	return os.WriteFile(filename, []byte(output), 0o644)
}
