package tests

import (
	"reflect"
	"testing"

	"go-reloaded/pipeline"
)

func TestApplyCaseTransform(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		expect []string
	}{
		{
			name:   "Basic capitalization",
			input:  []string{"hello", "world"},
			expect: []string{"Hello", "World"},
		},
		{
			name:   "Upper inside quotes",
			input:  []string{`She`, `said`, `"`, `hi`, `"`, `there`},
			expect: []string{"She", "Said", `"`, "HI", `"`, "There"},
		},
		{
			name:   "Lower inside quotes",
			input:  []string{`She`, `said`, `"`, `HI`, `"`, `there`},
			expect: []string{"She", "Said", `"`, "HI", `"`, "There"},
		},
		{
			name:   "Empty input",
			input:  []string{},
			expect: []string{},
		},
		{
			name:   "No quotes",
			input:  []string{"go", "lang"},
			expect: []string{"Go", "Lang"},
		},
		{
			name:   "Mixed quotes and normal words",
			input:  []string{`this`, `"`, `is`, `"`, `test`},
			expect: []string{"This", `"`, "IS", `"`, "Test"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pipeline.ApplyCaseTransform(tt.input)
			if !reflect.DeepEqual(result, tt.expect) {
				t.Errorf("ApplyCaseTransform() = %v, want %v", result, tt.expect)
			}
		})
	}
}
