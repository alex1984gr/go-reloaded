package tests

import (
	"strings"
	"testing"

	"go-reloaded/pipeline"
)

func TestReplaceBin(t *testing.T) {
	input := []string{"1010", "και", "110"}
	joined := strings.Join(input, " ")
	expected := "10 και 6"

	result := pipeline.ReplaceBin(joined)

	if result != expected {
		t.Errorf("Expected '%v', got '%v'", expected, result)
	}
}
func TestReplaceBin_NoBinary(t *testing.T) {
	input := []string{"Καλημέρα", "κόσμε"}
	joined := strings.Join(input, " ")
	expected := "Καλημέρα κόσμε"

	result := pipeline.ReplaceBin(joined)

	if result != expected {
		t.Errorf("Expected '%v', got '%v'", expected, result)
	}
}
func TestReplaceBin_Mixed(t *testing.T) {
	input := []string{"Το", "101", "είναι", "binary", "για", "5"}
	joined := strings.Join(input, " ")
	expected := "Το 5 είναι binary για 5"

	result := pipeline.ReplaceBin(joined)

	if result != expected {
		t.Errorf("Expected '%v', got '%v'", expected, result)
	}
}
