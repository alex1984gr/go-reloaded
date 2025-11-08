package TestFixQuotes_Simple

import (
	"go-reloade/pipeline"
	"reflect"
	"testing"
)

func TestFixArticles(t *testing.T) {
	test := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "A defore consonant stays A",
			input:    []string{"a", "banana"},
			expected: []string{"a", "banana"},
		},
		{
			name:     "A before vowel becomes An",
			input:    []string{"a", "apple"},
			expected: []string{"an", "apple"},
		},
		{
			name:     "An before consonant becomes A",
			input:    []string{"an", "dog"},
			expected: []string{"a", "dog"},
		},
		{
			name:     "silent H word gets An",
			input:    []string{"a", "hour"},
			expected: []string{"an", "hour"},
		},
		{
			name:     "Proper noun with capitalized A",
			input:    []string{"A", "apple"},
			expected: []string{"An", "apple"},
		},
		{
			name:     "Proper noun with capitalized An",
			input:    []string{"An", "banana"},
			expected: []string{"A", "banana"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pipeline.FixArticles(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Test %s failed.\nInput: %v\nExpected: %v\nGot: %v", tt.name, tt.input, tt.expected, result)
			}
		})
	}
}
