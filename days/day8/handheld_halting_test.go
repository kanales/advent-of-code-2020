package day8

import (
	"strings"
	"testing"
)

func TestParseProgram(t *testing.T) {
	input := []byte(strings.Join([]string{
		"nop +0",
		"acc +1",
		"jmp +4",
	}, "\n"))

	expect := []HandheldInstruction{
		{Op: "nop", Arg: 0},
		{Op: "acc", Arg: 1},
		{Op: "jmp", Arg: 4},
	}

	got := ParseHandheldProgram(input)
	for i := range expect {
		if got.Get(i) != expect[i] {
			t.Errorf("program.Get(%d) = %v; want %v", i, got.Get(i), expect[i])
		}
	}
}

func TestRunHandheldProgram(t *testing.T) {
	input := []byte(strings.Join([]string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}, "\n"))
	prog := ParseHandheldProgram(input)

	expect := 5
	got, ok := prog.Run()
	if ok || expect != got {
		t.Errorf("prog.Run(true) = %d, %v; want %d, %v", got, ok, expect, false)
	}
}

func TestPermute(t *testing.T) {
	input := []byte(strings.Join([]string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}, "\n"))

	prog := ParseHandheldProgram(input)
	prog.permute(0)
	got := prog.Get(0).Op
	expect := "jmp"
	if got != expect {
		t.Errorf("%q -> %q; want %q", "nop", got, expect)
	}
}

func TestRunFixProgram(t *testing.T) {
	input := []byte(strings.Join([]string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}, "\n"))
	prog := ParseHandheldProgram(input)

	expect := 7
	got, ok := prog.Fix()

	if !ok || expect != got {
		t.Errorf("prog.StaticAnalysis() = %d, %v; want %d, %v", got, ok, expect, true)
	}

	got, ok = prog.Run()
	expect = 8

	if !ok || expect != got {
		t.Errorf("prog.Run() = %d, %v; want %d, %v", got, ok, expect, true)
	}
}
