package day14

import (
	"reflect"
	"sort"
	"testing"
)

func TestParseMask(t *testing.T) {
	input := []byte("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")
	expect := Mask{Ones: 1 << 6, Zeros: 1 << 1}
	got := ParseMask(input)

	if got.Ones != expect.Ones || got.Zeros != expect.Zeros {
		t.Errorf("ParseMask(...) = %+v; want %+v", got, expect)
	}
}

func TestParseProgram(t *testing.T) {
	input := []byte("mask = X1011100000X111X01001000001110X00000\nmem[4616] = 8311689\nmem[8936] = 240\nmem[58007] = 369724\nmask = 10X0111X01X0XX110X10100X1001X000010X\nmem[41137] = 232605\nmem[33757] = 1437435\n")
	expect := Program{
		{
			Mask: ParseMask([]byte("X1011100000X111X01001000001110X00000")),
			Instructions: []Instruction{
				{Address: 4616, Value: 8311689},
				{Address: 8936, Value: 240},
				{Address: 58007, Value: 369724},
			},
		},
		{
			Mask: ParseMask([]byte("10X0111X01X0XX110X10100X1001X000010X")),
			Instructions: []Instruction{
				{Address: 41137, Value: 232605},
				{Address: 33757, Value: 1437435},
			},
		},
	}
	got := ParseProgram(input)

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("ParseProgram(...) = %+v;\n want %+v", got, expect)
	}
}

func TestMaskFloat(t *testing.T) {
	cases := []struct {
		Mask    Mask
		Address uint64
		Expect  []uint64
	}{
		{Mask: ParseMask([]byte("000000000000000000000000000000X1001X")), Address: 42, Expect: []uint64{
			26, 27, 58, 59,
		}},
		{Mask: ParseMask([]byte("00000000000000000000000000000000X0XX")), Address: 26, Expect: []uint64{
			16, 17, 18, 19, 24, 25, 26, 27,
		}},
	}

	for _, c := range cases {
		got := c.Mask.Float(c.Address)
		sort.Slice(got, func(i, j int) bool { return got[i] < got[j] })

		if !reflect.DeepEqual(c.Expect, got) {
			t.Errorf("%+v.Float(%d) = %v; want %v", c.Mask, c.Address, got, c.Expect)
		}
	}
}

func TestMaskApply(t *testing.T) {
	mask := Mask{Ones: 1 << 6, Zeros: 1 << 1}
	cases := []struct {
		Input  uint64
		Expect uint64
	}{
		{Input: 11, Expect: 73},
		{Input: 101, Expect: 101},
		{Input: 0, Expect: 64},
	}

	for _, c := range cases {
		got := mask.Apply(c.Input)
		if got != c.Expect {
			t.Errorf("mask.Apply(%d) = %d; want %d", c.Input, got, c.Expect)
		}
	}
}

func TestRun(t *testing.T) {
	input := Program{
		{
			Mask: Mask{Ones: 1 << 6, Zeros: 1 << 1},
			Instructions: []Instruction{
				{Address: 8, Value: 11},
				{Address: 7, Value: 101},
				{Address: 8, Value: 0},
			},
		},
		{
			Mask: Mask{Ones: 1 << 6, Zeros: 1 << 1},
			Instructions: []Instruction{
				{Address: 8, Value: 11},
				{Address: 7, Value: 101},
				{Address: 8, Value: 0},
			},
		},
	}
	expect := map[uint64]uint64{
		7: 101,
		8: 64,
	}

	got := input.Run(nil)

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("input.Run(nil) = %v; want %v", got, expect)
	}
}

func TestRunV2(t *testing.T) {
	input := Program{
		{
			Mask: ParseMask([]byte("000000000000000000000000000000X1001X")),
			Instructions: []Instruction{
				{Address: 42, Value: 100},
			},
		},
		{
			Mask: ParseMask([]byte("00000000000000000000000000000000X0XX")),
			Instructions: []Instruction{
				{Address: 26, Value: 1},
			},
		},
		{
			Mask: ParseMask([]byte("00000000000000000000000000000000X100")),
			Instructions: []Instruction{
				{Address: 8, Value: 127},
			},
		},
		{
			Mask: ParseMask([]byte("XX0000000000000000000000000011111111")),
			Instructions: []Instruction{
				{Address: 1, Value: 2020},
			},
		},
	}
	expect := map[uint64]uint64{
		//26: 100,
		//27: 100,
		58:                  100,
		59:                  100,
		16:                  1,
		17:                  1,
		18:                  1,
		19:                  1,
		24:                  1,
		25:                  1,
		26:                  1,
		27:                  1,
		4:                   127,
		12:                  127,
		(34359738368 + 255): 2020,
		(51539607552 + 255): 2020,
		(17179869184 + 255): 2020,
		255:                 2020,
	}

	// sum := func(m map[uint64]uint64) uint64 {
	// 	acc := uint64(0)
	// 	for _, x := range m {
	// 		acc += x
	// 	}
	// 	return acc
	// }

	got := input.RunV2(nil)

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("input.Run(nil) = \n%v\n; want \n%v", got, expect)
	}
}
