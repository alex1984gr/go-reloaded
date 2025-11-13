package tests // Test package for files in /tests folder

import (
	"fmt"     // For debug printing
	"reflect" // To compare slices
	"testing" // Go's testing framework

	"go-reloaded/pipeline" // Import the pipeline package containing ApplyCaseTransformations
)

// TestApplyCaseTransformations verifies uppercase, lowercase, and capitalize transformations
// including single-word and multi-word transformations.
func TestApplyCaseTransformations(t *testing.T) {
	tests := []struct {
		name     string   // Name of the test case
		input    []string // Input tokenized words
		expected []string // Expected output after transformations
	}{
		{
			name:     "Uppercase single word",
			input:    []string{"Ready,", "set,", "go", "(up)", "!"},
			expected: []string{"Ready,", "set,", "GO", "!"},
		},
		{
			name:     "Lowercase single word",
			input:    []string{"I", "should", "stop", "SHOUTING", "(low)"},
			expected: []string{"I", "should", "stop", "shouting"},
		},
		{
			name:     "Capitalize single word",
			input:    []string{"Welcome", "to", "the", "brooklyn", "bridge", "(cap)"},
			expected: []string{"Welcome", "to", "the", "brooklyn", "Bridge"}, // Fixed expected: last word before marker is capitalized
		},
		{
			name:     "Uppercase multiple words",
			input:    []string{"This", "is", "so", "exciting", "(up, 2)"},
			expected: []string{"This", "is", "SO", "EXCITING"},
		},
		{
			name:     "Lowercase multiple words",
			input:    []string{"PLEASE", "STOP", "YELLING", "(low, 3)"},
			expected: []string{"please", "stop", "yelling"},
		},
		{
			name:     "Capitalize multiple words",
			input:    []string{"welcome", "to", "brooklyn", "bridge", "(cap, 2)"},
			expected: []string{"welcome", "to", "Brooklyn", "Bridge"},
		},
		{
			name:     "No marker",
			input:    []string{"No", "changes", "here"},
			expected: []string{"No", "changes", "here"},
		},
		{
			name:     "Marker at beginning does nothing",
			input:    []string{"(up)", "hello"},
			expected: []string{"hello"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Debug: show test case name
			fmt.Printf("\n[DEBUG] Running test: %s\n", tt.name)
			// Debug: show input tokens
			fmt.Printf("[DEBUG] Input tokens: %v\n", tt.input)

			// Call the function being tested
			result := pipeline.ApplyCaseTransformations(tt.input)

			// Debug: show output tokens
			fmt.Printf("[DEBUG] Output tokens:   %v\n", result)
			// Debug: show expected tokens
			fmt.Printf("[DEBUG] Expected tokens: %v\n", tt.expected)

			// Compare result with expected
			if !reflect.DeepEqual(result, tt.expected) {
				// If not equal, fail the test
				t.Errorf("\n❌ Test failed: %s\nExpected: %v\nGot:      %v", tt.name, tt.expected, result)
			} else {
				// Otherwise, success message
				fmt.Printf("✅ Test passed: %s\n", tt.name)
			}
		})
	}
}
