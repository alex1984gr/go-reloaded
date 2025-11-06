package tests

import(
	"testing"
	"go-reloaded/pipeline"
)
func TestFixQuotes_Simple(t *testing.T){
	input := []string{'"', "Hello", "world", '"'}
	expected := []string{'"Hello World"'}
	result := pipeline.FixQuotes(input)

	if len(result) != len(expected) {
		t.Fatalf("Expected %d results, got %d", len(expected), len(result))
	}
	for i := range result {
		if result[i] != expected[i]{
			t.Errorf("Expected %v, got %v", expected[i], result[i])
		}
	}
}
func TestFixQuotes_Unmatched(t *testing.T){
	input := []string{"Hello",'"World"'}
	Expected := []string{"Hello", '"World"'}
	result := pipeline.FixQuotes(input)
	if len(result) != len(Expected) {
		t.Fatalf("Expected %d results, got %d", len(expected), len(result))
	}
	for i := range result {
		if result[i] != expected[i]{
			t.Errorf("Expected %v, got %v", expected[i], result[i])
		}
	}
}
func TestFixQuotes_NoQuotes(t *testing.T){
	input := []string{"Hello", "World"}
	expected := []string{"Hello", "World"}
	result := pipeline.FixQuotes(input)
	if len(result) != len(expected) {
		t.Fatalf("Expected %d results, got %d", len(expected), len(result))
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], result[i])
		}

	}
}	