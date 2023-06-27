package array

import (
	"fmt"
	"testing"
)

var testTable = []struct {
	input []int
	min   int
	max   int
}{
	{input: []int{}, min: -1, max: -1},
	{input: []int{1, 4, 5, 8, 2, 5, 8}, min: 1, max: 8},
	{input: []int{992, 124215, 122, 78, 12327, 1324}, min: 78, max: 124215},
	{input: []int{234, 1234, 6, 1, 7, 1245, 811, 6, 83}, min: 1, max: 1245},
}

func BenchmarkFindMinimum(b *testing.B) {
	for _, v := range testTable {
		b.Run(fmt.Sprintf("input: %v", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				got := findMinimum(v.input)
				if v.min != got {
					b.Errorf("Expected %v, but got %v", v.min, got)
				}
			}
		})
	}
}

func BenchmarkFindMaximum(b *testing.B) {
	for _, v := range testTable {
		b.Run(fmt.Sprintf("input: %v", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				got := findMaximum(v.input)
				if v.max != got {
					b.Errorf("Expected %v, but got %v", v.max, got)
				}
			}
		})
	}
}
