package utils

import "testing"

func TestGenerateSlug(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello World!", "hello-world"},
		{"A-Test-Case", "test-case"},
		{"#1 Ranking%", "sharp1-ranking-percent"},
		{"the-best example", "best-example"},
		{"Multiple   Spaces", "multiple-spaces"},
		{"Special@Characters!", "specialcharacters"},
		{"Numbers123", "numbers123"},
		{"a simple slug", "simple-slug"},
		{"Trailing - spaces ", "trailing-spaces"},
	}

	for _, test := range tests {
		output := GenerateSlug(test.input)
		if output != test.expected {
			t.Errorf("GenerateSlug(%q) = %q; want %q", test.input, output, test.expected)
		}
	}
}
