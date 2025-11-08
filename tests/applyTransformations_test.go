package tests

import (
	"go-reloaded/pipeline"
	"reflect"
	"testing"
)

func TestApplyTransfomations_Basic(t *testing.T) {
	input := []string{"a", "apple", "said", "\"hello\"", "world", "."}
	expected := []string{"an", "apple", "said", "\"Hello\"", "world", "."}
	result := pipeline.ApplyTransformations(input)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}
func TestApplyTransfomations_QuotesAndCase(t *testing.T) {
	input := []string{"she", "said", "\"hi\"", "there", "."}
	expected := []string{"she", "said", "\"HI\"", "there", "."}
	result := pipeline.ApplyTransformations(input)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}
func TestApplyTransfomations_PunctuationSpacing(t *testing.T) {
	input := []string{"Hello", ",", "world", "!"}
	expected := []string{"Hello", "world!"}
	result := pipeline.ApplyTransformation(input)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}
func TestApplyTransformations_EmptyInput(t *testing.T) {
	input := []string{}
	expected := []string{}
	result := pipeline.ApplyTransformation(input)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}
