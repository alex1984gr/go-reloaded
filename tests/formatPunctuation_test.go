package tests // Test package for /tests folder

import (
	"fmt"     // For debug printing
	"reflect" // To compare slices
	"testing" // Go's testing framework

	"go-reloaded/pipeline" // Import the pipeline package containing FormatPunctuation
)

// TestFormatPunctuation verifies that punctuation is formatted correctly
// according to the specified rules.
func TestFormatPunctuation(t *testing.T) {
	tests := []struct {
		name     string   // Name of the test case
		input    []string // Input tokenized words
		expected []string // Expected output after punctuation formatting
	}{
		{
			name:     "Basic punctuation spacing",
			input:    []string{"I", "was", "sitting", "over", "there", ",", "and", "then", "BAMM", "!!"},
			expected: []string{"I", "was", "sitting", "over", "there,", "and", "then", "BAMM!!"},
		},
		{
			name:     "Ellipsis spacing",
			input:    []string{"I", "was", "thinking", "...", "You", "were", "right"},
			expected: []string{"I", "was", "thinking...", "You", "were", "right"},
		},
		{
			name:     "Combined punctuation spacing",
			input:    []string{"Wait", "!?"},
			expected: []string{"Wait!?"},
		},
		{
			name:     "Single word quote",
			input:    []string{"I", "am", "exactly", "how", "they", "describe", "me", ":", "'", "awesome", "'"},
			expected: []string{"I", "am", "exactly", "how", "they", "describe", "me", ":", "'awesome'"},
		},
		{
			name:     "Multiple word quote",
			input:    []string{"As", "Elton", "John", "said", ":", "'", "I", "am", "the", "most", "well-known", "homosexual", "in", "the", "world", "'"},
			expected: []string{"As", "Elton", "John", "said", ":", "'I am the most well-known homosexual in the world'"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Debug: show test case name
			fmt.Printf("\n[DEBUG] Running test: %s\n", tt.name)
			// Debug: show input tokens
			fmt.Printf("[DEBUG] Input tokens: %v\n", tt.input)

			// Call the function being tested
			result := pipeline.FormatPunctuation(tt.input)

			// Debug: show output and expected
			fmt.Printf("[DEBUG] Output tokens:   %v\n", result)
			fmt.Printf("[DEBUG] Expected tokens: %v\n", tt.expected)

			// Compare result with expected
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("\n❌ Test failed: %s\nExpected: %v\nGot:      %v", tt.name, tt.expected, result)
			} else {
				fmt.Printf("✅ Test passed: %s\n", tt.name)
			}
		})
	}
}
