package tests

import (
	"go-reloaded/pipeline"
	"os"
	"testing"
)

func TestReadInput(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "input_*.txt") // Create a temporary file for the test
	if err != nil {
		t.Fatalf("❌ Failed to create temporary file: %v", err) // Delete the file after the test finishes
	}
	defer os.Remove(tmpFile.Name())

	// Write simple content into it for testing
	content := "Hello, Go Reloaded!"
	tmpFile.WriteString(content)
	tmpFile.Close()

	// Call the readInput function from the pipeline
	result, err := pipeline.ReadInput(tmpFile.Name())

	// Check if an error was returned (there shouldn't be one)
	if err != nil {
		t.Fatalf("❌ Unexpected error: %v", err)
	}

	// Verify that the result matches the original content
	if result != content {
		t.Errorf("❌ Expected '%s', receive '%s'", content, result)
	}

	// If execution reaches this point, everything went well
	t.Logf("✅ The readInput read the file correctly : %s", tmpFile.Name())
}

// TestReadInput_FileNotFound checks the behavior when the file does not exist
func TestReadInput_FileNotFound(t *testing.T) {
	// Call the function with a non-existent path
	_, err := pipeline.ReadInput("non_existing_file.txt")

	// If it doesn't return an error, the test fails
	if err == nil {
		t.Errorf("❌ An error was expected for a non-existent file, but none was received.")
	} else {
		t.Logf("✅ the non_existent file handling works correctly: %v", err)
	}
}
