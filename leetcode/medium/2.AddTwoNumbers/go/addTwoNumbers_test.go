package addtwonumbers

import (
	"reflect"
	"testing"

)

const (
	ColorGreen = "\033[32m"
	ColorRed   = "\033[31m"
	ColorReset = "\033[0m"
)

// Helper: Convert a slice of ints to a Linked List
func sliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	curr := head
	for i := 1; i < len(nums); i++ {
		curr.Next = &ListNode{Val: nums[i]}
		curr.Next.Next = nil // strictly for clarity
		curr = curr.Next
	}
	return head
}

// Helper: Convert a Linked List back to a slice for easy comparison
func listToSlice(l *ListNode) []int {
	var nums []int
	for l != nil {
		nums = append(nums, l.Val)
		l = l.Next
	}
	return nums
}

func TestAddTwoNumbers(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string
		l1       []int
		l2       []int
		expected []int
	}{
		{
			name:     "Standard Case: 342 + 465 = 807",
			l1:       []int{2, 4, 3},
			l2:       []int{5, 6, 4},
			expected: []int{7, 0, 8},
		},
		{
			name:     "Zeroes: 0 + 0 = 0",
			l1:       []int{0},
			l2:       []int{0},
			expected: []int{0},
		},
		{
			name:     "Different Lengths: 99 + 1 = 100",
			l1:       []int{9, 9},
			l2:       []int{1},
			expected: []int{0, 0, 1},
		},
		{
			name:     "Carry at the end: 9+9 = 18",
			l1:       []int{9},
			l2:       []int{9},
			expected: []int{8, 1},
		},
		{
			name:     "Long Carry: 9999 + 1 = 10000",
			l1:       []int{9, 9, 9, 9},
			l2:       []int{1},
			expected: []int{0, 0, 0, 0, 1},
		},
		{
			name:     "Big Integers (Overflow test)",
			l1:       []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			l2:       []int{4, 6, 5},
			expected: []int{5, 6, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		},
		{
			name:			"Really Big Integers",
			l1: 			[]int{6,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1},
			l2:				[]int{5,6,4},
			expected:	[]int{1,7,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1},
		},
		{
			name:     "Empty-style leading zeros (though problem says non-empty)",
			l1:       []int{0, 1}, // represents 10
			l2:       []int{0, 2}, // represents 20
			expected: []int{0, 3},
		},
		{
			name:     "All 9s",
			l1: 			[]int{9,9,9,9,9,9,9},
			l2: 			[]int{9,9,9,9},
			expected: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert input slices to Linked Lists
			l1 := sliceToList(tt.l1)
			l2 := sliceToList(tt.l2)

			// Call your function
			// resultList := AddTwoNumbersSlow(l1, l2)
			// resultList := AddTwoNumbersSlow(l1, l2)
			resultList := AddTwoNumbersFast2(l1, l2)

			// Convert result back to slice to compare
			actual := listToSlice(resultList)

			// DeepEqual compares slice contents
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("%sFAIL: AddTwoNumbers(%v, %v) Result: %v, Expected: %v%s",ColorRed, tt.l1, tt.l2, actual, tt.expected, ColorReset)
			} else {
				t.Logf("%sPASS: %s%s", ColorGreen, tt.name, ColorReset)
			}
		})
	}
}