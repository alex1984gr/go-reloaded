package pipeline

import (
	"os"
	"strings"
)

func JoinTokens(tokens []string) string {
	return strings.Join(tokens, "")
}

func WriteOutput(filename string, input []string) error {
	output := strings.Join(input, " ")
	return os.WriteFile(filename, []byte(output), 0o644)
}
