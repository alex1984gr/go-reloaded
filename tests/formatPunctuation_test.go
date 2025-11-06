package tests

import (
	"testing"

	"go-reloaded/pipeline"
)

func TestFormatPunctuation(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "simple punctuation",
			input:    []string{"Hello", ",", "world", "!"},
			expected: []string{"Hello,", "world!"},
		},
		{
			name:     "no punctuation",
			input:    []string{"Just", "words"},
			expected: []string{"Just", "words"},
		},
		{
			name:     "space before punctuation",
			input:    []string{"Wow", " ", "!"},
			expected: []string{"Wow!"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pipeline.FormatPunctuation(tt.input)

			if len(result) != len(tt.expected) {
				t.Fatalf("Expected length %d, got %d", len(tt.expected), len(result))
			}

			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Expected %q at index %d, got %q", tt.expected[i], i, result[i])
				}
			}
		})
	}
}
