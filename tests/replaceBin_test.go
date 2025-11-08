package tests

import (
	"strings"
	"testing"

	"go-reloaded/pipeline"
)

func TestReplaceBin(t *testing.T) {
	input := []string{"1010", "και", "110"}
	expected := []string{"10", "και", "6"}

	result := pipeline.ReplaceBin(input)

	if strings.Join(result, " ") != strings.Join(expected, " ") {
		t.Errorf("Expected '%v', got '%v'", strings.Join(expected, " "), strings.Join(result, " "))
	}
}

func TestReplaceBin_NoBinary(t *testing.T) {
	input := []string{"Καλημέρα", "κόσμε"}
	expected := []string{"Καλημέρα", "κόσμε"}

	result := pipeline.ReplaceBin(input)

	if strings.Join(result, " ") != strings.Join(expected, " ") {
		t.Errorf("Expected '%v', got '%v'", strings.Join(expected, " "), strings.Join(result, " "))
	}
}

func TestReplaceBin_Mixed(t *testing.T) {
	input := []string{"Το", "101", "είναι", "binary", "για", "5"}
	expected := []string{"Το", "5", "είναι", "binary", "για", "5"}

	result := pipeline.ReplaceBin(input)

	if strings.Join(result, " ") != strings.Join(expected, " ") {
		t.Errorf("Expected '%v', got '%v'", strings.Join(expected, " "), strings.Join(result, " "))
	}
}
