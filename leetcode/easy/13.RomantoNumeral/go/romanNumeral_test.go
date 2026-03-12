package romannumeral

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"Single Digit", "I", 1},
		{"Simple Addition", "III", 3},
		{"Mixed Addition", "LVIII", 58},
		{"Subtractive IV", "IV", 4},
		{"Subtractive IX", "IX", 9},
		{"Subtractive XL", "XL", 40},
		{"Subtractive XC", "XC", 90},
		{"Subtractive CD", "CD", 400},
		{"Subtractive CM", "CM", 900},
		{"Complex Case 1994", "MCMXCIV", 1994},
		{"Complex Case 3999", "MMMCMXCIX", 3999},
		{"Year 2026", "MMXXVI", 2026},
		{"Mid Range", "DCLXVI", 666},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RomanToInt(tt.input)
			if result != tt.expected {
				t.Errorf("RomanToInt(%s) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}