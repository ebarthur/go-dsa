package strings

import (
	"testing"
)

func TestJohnVenn(t *testing.T) {
	tests := []struct {
		input, expected string
		errExpected     bool
	}{
		{"", "", true},                                        // Invalid input (empty string)
		{"1 2 3\n4 5 6", "", false},                           // No common elements
		{"1 2 3\n3 2 1", "1 2 3", false},                      // All elements are common
		{"A B C\nC B A", "A B C", false},                      // All elements are common (letters)
		{"1 2 3 A B C\nX 11 G M 2 9 3 C N R", "2 3 C", false}, // Mixed input
		{"1 2 3\n", "", true},                                 // Invalid input (only one line)
	}

	for i, test := range tests {
		result, err := JohnVenn(test.input)
		
		if test.errExpected {
			if err == nil {
				t.Fatalf("Failed test case #%d. Expected an error but got none", i)
			}
		} else {
			if err != nil {
				t.Fatalf("Failed test case #%d. Unexpected error: %v", i, err)
			}
			if result != test.expected {
				t.Fatalf("Failed test case #%d. Expected %q but got %q", i, test.expected, result)
			}
		}
	}
}
