package day15

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := []byte("1,0,16,5,17,4\n")
	expect := []uint{1, 0, 16, 5, 17, 4}
	got := ParseInput(input)

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("ParseInput(%q) = %v; want %v", input, got, expect)
	}
}

func TestMemoryGame(t *testing.T) {
	inputs := [][]uint{{0, 3, 6}, {1, 3, 2}, {2, 1, 3}, {1, 2, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}, {1, 2, 3, 4}, {0, 5, 4, 1, 10, 14, 7}, {1, 0, 16, 5, 17, 4}}
	expects := []uint{436, 1, 10, 27, 78, 438, 1836, 10, 203, 0}
	rounds := uint(2020)
	for i := range inputs {
		input := inputs[i]
		expect := expects[i]
		got := MemoryGame(input, rounds)

		if expect != got {
			t.Errorf("MemoryGame(%v, %d) = %d; want %d", input, rounds, got, expect)
		}
	}

}

func BenchmarkMemoryGame(b *testing.B) {
	cases := [][]uint{{0, 5, 4, 1, 10, 14, 7}, {1, 0, 16, 5, 17, 4}}
	for _, c := range cases {
		b.Run(fmt.Sprintln("Benchmark", c), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MemoryGame(c, 30000000)
			}
		})
	}

}
