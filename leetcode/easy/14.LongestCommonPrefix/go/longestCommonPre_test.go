package longestcommonprefix

import "testing"

// func TestLongestCommonPrefix(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		strs     []string
// 		expected string
// 	}{
// 		{
// 			name:     "Example 1: Basic common prefix",
// 			strs:     []string{"flower", "flow", "flight"},
// 			expected: "fl",
// 		},
// 		{
// 			name:     "Example 2: No common prefix",
// 			strs:     []string{"dog", "racecar", "car"},
// 			expected: "",
// 		},
// 		{
// 			name:     "All identical strings",
// 			strs:     []string{"interspecies", "interspecies", "interspecies"},
// 			expected: "interspecies",
// 		},
// 		{
// 			name:     "Prefix is one whole string",
// 			strs:     []string{"apple", "app", "april"},
// 			expected: "ap",
// 		},
// 		{
// 			name:     "Single string in array",
// 			strs:     []string{"standalone"},
// 			expected: "standalone",
// 		},
// 		{
// 			name:     "One string is empty",
// 			strs:     []string{"", "b"},
// 			expected: "",
// 		},
// 		{
// 			name:     "Differing characters at first index",
// 			strs:     []string{"ab", "cb"},
// 			expected: "",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			result := longestCommonPrefix(tt.strs)
// 			if result != tt.expected {
// 				t.Errorf("longestCommonPrefix(%v) = %q; want %q", tt.strs, result, tt.expected)
// 			}
// 		})
// 	}
// }

func TestLongestCommonPrefix2(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected string
	}{
		{
			name:     "Example 1: Basic common prefix",
			strs:     []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			name:     "Example 2: No common prefix",
			strs:     []string{"dog", "racecar", "car"},
			expected: "",
		},
		{
			name:     "All identical strings",
			strs:     []string{"interspecies", "interspecies", "interspecies"},
			expected: "interspecies",
		},
		{
			name:     "Prefix is one whole string",
			strs:     []string{"apple", "app", "april"},
			expected: "ap",
		},
		{
			name:     "Single string in array",
			strs:     []string{"standalone"},
			expected: "standalone",
		},
		{
			name:     "One string is empty",
			strs:     []string{"", "b"},
			expected: "",
		},
		{
			name:     "Differing characters at first index",
			strs:     []string{"ab", "cb"},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestCommonPrefix2(tt.strs)
			if result != tt.expected {
				t.Errorf("longestCommonPrefix(%v) = %q; want %q", tt.strs, result, tt.expected)
			}
		})
	}
}