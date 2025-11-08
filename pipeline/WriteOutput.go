package pipeline

import (
	"os"
	"strings"
)

func WriteOutPut(filename string, input []string) error {
	output := strings.Join(input, " ")
	return os.WriteFile(filename, []byte(output), 0644)
}
