package palindrome

import "testing"

// func TestIsPalindromeSlow(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		input    int
// 		expected bool
// 	}{
// 		{"Example 1: Positive Palindrome", 121, true},
// 		{"Example 2: Negative Number", -121, false},
// 		{"Example 3: Ends in Zero", 10, false},
// 		{"Single Digit", 7, true},
// 		{"Large Palindrome", 123454321, true},
// 		{"Zero", 0, true},
// 		{"Max Int32 Range", 2147483647, false},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			result := IsPalindromeSlow(tt.input)
// 			if result != tt.expected {
// 				t.Errorf("IsPalindrome(%d) = %v; want %v", tt.input, result, tt.expected)
// 			}
// 		})
// 	}
// }

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Example 1: Positive Palindrome", 121, true},
		{"Example 2: Negative Number", -121, false},
		{"Example 3: Ends in Zero", 10, false},
		{"Single Digit", 7, true},
		{"Large Palindrome", 123454321, true},
		{"Large Fake Palindrome", 123453321, false},
		{"Zero", 0, true},
		{"Max Int32 Range", 2147483647, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("IsPalindrome(%d) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

