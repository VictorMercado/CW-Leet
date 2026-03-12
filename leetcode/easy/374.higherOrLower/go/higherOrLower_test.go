package higherorlower

import "testing"

func TestGuessNumber(t *testing.T) {
	tests := []struct {
		name string
		n    int // The range 1..n
		pick int // The number the "game" picked
	}{
		{
			name: "Example 1: Standard range",
			n:    100,
			pick: 52,
		},
		{
			name: "Example 2: Minimum range",
			n:    1,
			pick: 1,
		},
		{
			name: "Example 3: Small range",
			n:    2,
			pick: 1,
		},
		{
			name: "Large Range: Max Int32",
			n:    2147483647, // 2^31 - 1
			pick: 1000000000,
		},
		{
			name: "Pick is the boundary (n)",
			n:    100,
			pick: 99,
		},
		{
			name: "Pick is the boundary (1)",
			n:    100,
			pick: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set the global mock target for this specific test
			Target = tt.pick
			
			// Assuming your solution function is named 'guessNumber'
			result := GuessNumber(tt.n)

			if result != tt.pick {
				t.Errorf("guessNumber(%d) with pick %d failed: got %d, want %d", tt.n, tt.pick, result, tt.pick)
			}
		})
	}
}