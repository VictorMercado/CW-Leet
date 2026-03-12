package twosum

import (
	"reflect"
	// "runtime"
	"math/rand"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		target int
		nums []int
		expected []int
	}{
		{ 9, []int{2,7,11,15}, []int{0,1} },
		{ 6, []int{3,2,4}, []int{1,2} },
		{ 6, []int{3,3}, []int{0,1} },
		{10, []int{5, 5}, []int{0, 1}},
		// Negative numbers
		{0, []int{-3, 4, 3, 90}, []int{0, 2}},
		// Mixed positive/negative
		{-3, []int{-1, -2, -3, -4, -5}, []int{0, 1}},
		// Target achievable by last two elements
		{8, []int{1, 2, 3, 5, 8, 3}, []int{2, 3}},
		// Large numbers
		{100000000, []int{99999999, 1}, []int{0, 1}},
		// Duplicates where only specific pair works
		{10, []int{1, 5, 5, 3, 7}, []int{1, 2}},
		// First and last element
		{8, []int{1, 4, 6, 7}, []int{0, 3}},
		// No valid pair (optional — may return nil or empty)
		{6, []int{1, 2, 3}, []int{}},
		// Same number used twice only if repeated
		{10, []int{5, 3, 5}, []int{0, 2}},
	}

	for _, tt := range tests {
		// var m1, m2 runtime.MemStats

		// runtime.ReadMemStats(&m1)
		result := TwoSum(tt.nums, tt.target)
		// runtime.ReadMemStats(&m2)
		// t.Logf("Memory used: %v KB\n", (m2.Alloc-m1.Alloc)/1024)

		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("target: %d slices not equal: got %v, want %v", tt.target, result, tt.expected)
		}
	}
}

func TestEvenMorePerfTwoSum(t *testing.T) {
	tests := []struct {
		target int
		nums []int
		expected []int
	}{
		{ 9, []int{2,7,11,15}, []int{0,1} },
		{ 6, []int{3,2,4}, []int{1,2} },
		{ 6, []int{3,3}, []int{0,1} },
		{10, []int{5, 5}, []int{0, 1}},
		// Negative numbers
		{0, []int{-3, 4, 3, 90}, []int{0, 2}},
		// Mixed positive/negative
		{-3, []int{-1, -2, -3, -4, -5}, []int{0, 1}},
		// Target achievable by last two elements
		{8, []int{1, 2, 3, 5, 8, 3}, []int{2, 3}},
		// Large numbers
		{100000000, []int{99999999, 1}, []int{0, 1}},
		// Duplicates where only specific pair works
		{10, []int{1, 5, 5, 3, 7}, []int{1, 2}},
		// First and last element
		{8, []int{1, 4, 6, 7}, []int{0, 3}},
		// No valid pair (optional — may return nil or empty)
		{6, []int{1, 2, 3}, []int{}},
		// Same number used twice only if repeated
		{10, []int{5, 3, 5}, []int{0, 2}},
	}

	for _, tt := range tests {
		// var m1, m2 runtime.MemStats

		// runtime.ReadMemStats(&m1)
		result := EvenMorePerfTwoSum(tt.nums, tt.target)
		// runtime.ReadMemStats(&m2)
		// t.Logf("Memory used: %v KB\n", (m2.Alloc-m1.Alloc)/1024)

		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("target: %d slices not equal: got %v, want %v", tt.target, result, tt.expected)
		}
	}
}

func generateLargeSlice(n int) []int {
    nums := make([]int, n)
    for i := 0; i < n; i++ {
        nums[i] = rand.Intn(n * 2) // random numbers
    }
    return nums
}

func BenchmarkTwoSumLarge(b *testing.B) {
    nums := generateLargeSlice(10000)
    target := nums[0] + nums[9999]

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        TwoSum(nums, target)
    }
}

func BenchmarkMorePerfTwoSumLarge(b *testing.B) {
    nums := generateLargeSlice(10000)
    target := nums[0] + nums[9999]

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        MorePerfTwoSum(nums, target)
    }
}

func BenchmarkEvenMorePerfTwoSumLarge(b *testing.B) {
    nums := generateLargeSlice(10000)
    target := nums[0] + nums[9999]

		b.ResetTimer()
    for i := 0; i < b.N; i++ {
        EvenMorePerfTwoSum(nums, target)
    }
}