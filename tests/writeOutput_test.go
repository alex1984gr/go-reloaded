package tests

import (
	"go-reloaded/pipeline"
	"os"
	"strings"
	"testing"
)

func TestWriteOutPut(t *testing.T) {
	input := []string{"Hello", "deautiful", "world!"}
	tmpFile := "test_output.txt"
	err := pipeline.WriteOutPut(tmpFile, input)
	if err != nil {
		t.Fatalf("WriteOutPut returned error: %v", err)
	}
	data, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}
	content := strings.TrimSpace(string(data))
	expected := "Hello beautiful world!"
	if content != expected {
		t.Errorf("Expected '%s', got '%s'", expected, content)
	}
	os.Remove(tmpFile)
}
func TestWriteOutput_EmptyInput(t *testing.T) {
	tmpFile := "test_output_empty.txt"
	input := []string{}

	err := pipeline.WriteOutPut(tmpFile, input)
	if err != nil {
		t.Errorf("Expected empty file, got '%s'", content)
	}
	os.Remove(tmpFile)
}
