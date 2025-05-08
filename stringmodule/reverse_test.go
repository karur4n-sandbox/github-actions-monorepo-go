package stringmodule_test

import (
	"testing"

	"github.com/karur4n-sandbox/github-actions-monorepo-go/stringmodule"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "reverse of 'hello' is 'olleh'",
			input:    "hello",
			expected: "olleh",
		},
		{
			name:     "reverse of 'world' is 'dlrow'",
			input:    "world",
			expected: "dlrow",
		},
		{
			name:     "reverse of 'golang' is 'gnalog'",
			input:    "golang",
			expected: "gnalog",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := stringmodule.Reverse(test.input)
			if result != test.expected {
				t.Errorf("Reverse(%q) = %q; expected %q", test.input, result, test.expected)
			}
		})
	}
}

func TestReverseHello(t *testing.T) {
	expected := "olleH"
	result := stringmodule.ReverseHello()
	if result != expected {
		t.Errorf("ReverseHello() = %q; expected %q", result, expected)
	}
}
