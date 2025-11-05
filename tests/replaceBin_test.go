package tests

import (
	"go-reloaded/pipeline"
	"testing"
)

func TestReplaceBin_ValidBinary(t *testing.T) {
	input := []string{"The", "value", "is", "10", "(bin)"}
	expected := []string{"The", "value", "is", "2"}
	result := pipeline.ReplaceBin(input)

	if len(result) != len(expected) {
		t.Fatalf("Expected length%d, got %d", len(expected), len(result))
	}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], result[i])
		}
	}
}
func TestReblaceBin_InvalidBinary(t *testing.T) {
	input := []string{"This", "is", "not", "a", "2", "(bin)"}
	expected := []string{"This", "is", "not", "a", "2"}
	result := pipeline.ReplaceBin(input)

	if len(result) != len(expected) {
		t.Fatalf("Expected %d, got %d", len(expected), len(result))
	}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], result[i])
		}
	}
}
func TestReplaceBIn_NoBinTag(t *testing.T) {
	input := []string{"Hello", "world"}
	expected := []string{"Hello", "world"}
	result := pipeline.ReplaceBin(input)

	if len(result) != len(expected) {
		t.Fatalf("Expected length %d, got %d", len(expected), len(result))
	}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], result[i])
		}
	}
}
