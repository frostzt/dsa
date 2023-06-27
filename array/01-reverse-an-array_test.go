package main

import (
	utilities "dsa/_utilities"
	"fmt"
	"testing"
)

// `cd` into this directory and run the following for results
// go test -bench=. -benchtime=1x -benchmem
// b.N will be set to 100 times for number of iterations

var integerTable = []struct {
	input []int
	want  []int
}{
	{input: []int{1, 2, 3, 4, 5}, want: []int{5, 4, 3, 2, 1}},
	{input: []int{99, 78, 83, 578, 490}, want: []int{490, 578, 83, 78, 99}},
	{input: []int{249, 532, 58, 23, 243908}, want: []int{243908, 23, 58, 532, 249}},
}

var stringsTable = []struct {
	input string
	want  string
}{
	{input: "Hello world", want: "dlrow olleH"},
	{input: "Sourav", want: "varuoS"},
	{input: "FizzBuzz", want: "zzuBzziF"},
}

func BenchmarkReverseIntegerArray(b *testing.B) {
	for _, v := range integerTable {
		b.Run(fmt.Sprintf("input: %d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				got := reverse(v.input)
				if _got, ok := got.([]int); ok {
					check := utilities.CompareSliceValueToValue(_got, v.want)
					if !check {
						b.Errorf("Expected %v, but got %v", v.want, got)
					}
				} else {
					b.Error("Something went wrong with types....")
				}
			}
		})
	}
}

func BenchmarkReverseStrings(b *testing.B) {
	for _, v := range stringsTable {
		b.Run(fmt.Sprintf("input: %s", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				got := reverse(v.input)
				if v.want != got {
					b.Errorf("Expected %s, but got %s", v.want, got)
				}
			}
		})
	}
}
