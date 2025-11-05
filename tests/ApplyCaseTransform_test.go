package tests

import (
	"go-reloaded/pipeline"
	"testing"
)

func TestApllyCaseTransform_upper(t *testing.T) {
	input := []string{"Hello", "world"}
	expected := []string{"HELLO", "WORLD"}
	result := pipeline.ApplyCaseTransform(input, "upper")
	if len(result) != len(expected) {
		t.Fatalf("Expected %d results, got %d", len(expected), len(result))
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected '%v', got '%v'", expected[i], result[i])
		}
	}
}
func TestApllyCaseTransform_Capitalize(t *testing.T) {
	input := []string{"hello", "world"}
	expected := []string{"Hello", "World"}
	result := pipeline.ApplyCaseTransform(input, "capitalize")
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected '%v', got '%v'", expected[i], result[i])
		}
	}
}
func TestApllyCaseTransform_Mixed(t *testing.T) {
	input := []string{"TeSt", "WoRd"}
	expected := []string{"TeSt", "WoRd"}
	result := pipeline.ApplyCaseTransform(input, "unknown")
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected '%v', got '%v'", expected[i], result[i])
		}
	}
}
