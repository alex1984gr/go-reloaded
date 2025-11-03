package tests

import (
	"go-reloaded/pipeline"
	"testing"
)

// TestReplaceHex_Basic checks simple hex-to-decimal conversions
func TestReplaceHex_Basic(t *testing.T) {
	input := []string{"The value is 1E (hex) and another is FF (hex)."}
	expected := []string{"The value is 30 and another is 255."}

	result := pipeline.ReplaceHex(input)

	if len(result) != len(expected) {
		t.Fatalf("Expected %d lines, got %d", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Line %d mismatch:\nExpected: %s\nGot: %s", i+1, expected[i], result[i])
		}
	}
}

// TestReplaceHex_Invalid checks handling of invalid hex patterns
func TestReplaceHex_Invalid(t *testing.T) {
	input := []string{"This is not a valid number (hex)", "G1 (hex) should stay the same."}
	expected := []string{"This is not a valid number (hex)", "G1 (hex) should stay the same."}

	result := pipeline.ReplaceHex(input)

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Invalid case failed at line %d:\nExpected: %s\nGot: %s", i+1, expected[i], result[i])
		}
	}
}

// TestReplaceHex_Multiple checks multiple conversions in one line
func TestReplaceHex_Multiple(t *testing.T) {
	input := []string{"Values: A (hex), 10 (hex), and 1F (hex)."}
	expected := []string{"Values: 10, 16, and 31."}

	result := pipeline.ReplaceHex(input)

	if len(result) != len(expected) {
		t.Fatalf("Expected %d lines, got %d", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Mismatch on line %d:\nExpected: %s\nGot: %s", i+1, expected[i], result[i])
		}
	}
}
