package mathmodule_test

import (
	"testing"

	"github.com/karur4n-sandbox/github-actions-monorepo-go/mathmodule"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "1 + 1 = 2",
			input:    1,
			expected: 2,
		},
		{
			name:     "2 + 1 = 3",
			input:    2,
			expected: 3,
		},
		{
			name:     "3 + 1 = 4",
			input:    3,
			expected: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := mathmodule.PlusOne(test.input)
			if result != test.expected {
				t.Errorf("PlusOne(%d) = %d; expected %d", test.input, result, test.expected)
			}
		})
	}
}
