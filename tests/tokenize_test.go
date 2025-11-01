package tests

import (
	"go-reloaded/pipeline"
	"reflect"
	"testing"
)

func TestTokenize_SimpleWords(t *testing.T) {
	input := "Hello world"
	expected := []string{"Hello", "world"}

	result := pipeline.Tokenize(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Tokenize failed. Expected %v, got %v", expected, result)
	}
}
func TestTokenize_Punctuation(t *testing.T) {
	input := "Hello, world!"
	expected := []string{"Hello", ",", "world", "!"}

	result := pipeline.Tokenize(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Tokenize failed. Expected %v, got %v", expected, result)
	}
}

func TestTokenize_MixedTags(t *testing.T) {
	input := "Go (up) to the 10 (hex) level"
	expected := []string{"Go", "(up)", "to", "the", "10", "(hex)", "level"}

	result := pipeline.Tokenize(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Tokenize failed. Expected %v, got %v", expected, result)
	}
}

func TestTokenize_EmptyString(t *testing.T) {
	input := ""
	expected := []string{}

	result := pipeline.Tokenize(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Tokenize failed. Expected empty slice, got %v", result)
	}
}
