package tests

import (
	"reflect"
	"testing"

	"go-reloaded/pipeline"
)

// TestTokenize_Basic ελέγχει τη βασική λειτουργία της tokenize
func TestTokenize_Basic(t *testing.T) {
	input := []rune(`Hello <world> "Go"`)
	expected := []string{"Hello", "<", "world", ">", "\"", "Go", "\""}

	result := pipeline.Tokenize(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestTokenize_WithSpaces ελέγχει ότι αγνοούνται πολλαπλά κενά
func TestTokenize_WithSpaces(t *testing.T) {
	input := []rune("Hello   world")
	expected := []string{"Hello", "world"}

	result := pipeline.Tokenize(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestTokenize_OnlySpecials ελέγχει τη συμπεριφορά μόνο με ειδικούς χαρακτήρες
func TestTokenize_OnlySpecials(t *testing.T) {
	input := []rune(`<>""`)
	expected := []string{"<", ">", "\"", "\""}

	result := pipeline.Tokenize(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestTokenize_EmptyInput ελέγχει τι γίνεται με κενό input
func TestTokenize_EmptyInput(t *testing.T) {
	input := []rune("")
	expected := []string{}

	result := pipeline.Tokenize(input)

	if len(result) != len(expected) {
		t.Errorf("Expected %v tokens, got %v", len(expected), len(result))
	}

}
