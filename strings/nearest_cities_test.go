package strings

import (
	"reflect"
	"testing"
)

func TestNearestCities(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []int32
		hasError bool
	}{
		{
			name:     "Valid input with multiple cities",
			input:    "Accra, 121; Kumasi, 349; Takoradi, 718",
			expected: []int32{121, 228, 369}, // 121, 349-121, 718-349
			hasError: false,
		},
		{
			name:     "Single city",
			input:    "Accra, 100",
			expected: []int32{100},
			hasError: false,
		},
		{
			name:     "Two cities",
			input:    "City1, 50; City2, 150",
			expected: []int32{50, 100}, // 50, 150-50
			hasError: false,
		},
		{
			name:     "Cities in reverse order",
			input:    "Takoradi, 718; Kumasi, 349; Accra, 121",
			expected: []int32{121, 228, 369}, // Should sort: 121, 349-121, 718-349
			hasError: false,
		},
		{
			name:     "Zero distances",
			input:    "City1, 0; City2, 100; City3, 50",
			expected: []int32{0, 50, 50}, // Sorted: 0, 50, 100 -> 0, 50-0, 100-50
			hasError: false,
		},
		{
			name:     "Negative distances",
			input:    "City1, -10; City2, 20; City3, 5",
			expected: []int32{-10, 15, 15}, // Sorted: -10, 5, 20 -> -10, 5-(-10), 20-5
			hasError: false,
		},
		{
			name:     "Empty input",
			input:    "",
			expected: nil,
			hasError: true,
		},
		{
			name:     "Only semicolons",
			input:    ";;;",
			expected: nil,
			hasError: true,
		},
		{
			name:     "Invalid format - no comma",
			input:    "Accra 121; Kumasi 349",
			expected: nil,
			hasError: true,
		},
		{
			name:     "Invalid format - non-numeric distance",
			input:    "Accra, abc; Kumasi, 349",
			expected: []int32{349},
			hasError: false,
		},
		{
			name:     "Mixed valid and invalid entries",
			input:    "Accra, 121; Invalid Entry; Kumasi, 349; Bad, xyz; Takoradi, 718",
			expected: []int32{121, 228, 369},
			hasError: false,
		},
		{
			name:     "Whitespace handling",
			input:    " Accra , 121 ; Kumasi , 349 ; Takoradi , 718 ",
			expected: []int32{121, 228, 369},
			hasError: false,
		},
		{
			name:     "Duplicate distances",
			input:    "City1, 100; City2, 100; City3, 200",
			expected: []int32{100, 0, 100}, // Sorted: 100, 100, 200 -> 100, 0, 100
			hasError: false,
		},
		{
			name:     "Large numbers",
			input:    "City1, 2147483647; City2, 1000000000",
			expected: []int32{1000000000, 1147483647}, // Sorted order
			hasError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NearestCities(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestAbsFunction(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int32
		expected int32
	}{
		{"Positive difference", 10, 5, 5},
		{"Negative difference", 5, 10, 5},
		{"Zero difference", 5, 5, 0},
		{"Negative numbers", -5, -10, 5},
		{"Mixed signs", -5, 10, 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := abs(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("abs(%d, %d) = %d, expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
