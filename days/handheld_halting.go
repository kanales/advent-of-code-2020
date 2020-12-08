package days

import (
	"bytes"
	"strconv"
)

type HandheldInstruction struct {
	Op  string
	Arg int
}
type HandheldProgram []HandheldInstruction

func (prog *HandheldProgram) Get(i int) HandheldInstruction {
	return (*prog)[i]
}

func (prog *HandheldProgram) Len() int {
	return len(*prog)
}

func ParseHandheldProgram(input []byte) HandheldProgram {
	lines := bytes.Split(bytes.TrimRight(input, "\n"), NL)
	program := make(HandheldProgram, len(lines))
	for i := range lines {
		split := bytes.Split(lines[i], []byte{' '})
		arg, err := strconv.Atoi(string(split[1]))
		if err != nil {
			panic(err)
		}
		program[i] = HandheldInstruction{
			Op:  string(split[0]),
			Arg: arg,
		}
	}
	return program
}

// Run returns the result of an accum and wether it finished normally
func (prog *HandheldProgram) Run() (int, bool) {
	accum := 0
	pc := 0
	memory := make(map[int]bool)
	for pc < len(*prog) {
		if memory[pc] {
			return accum, false
		}

		memory[pc] = true
		instruction := prog.Get(pc)

		switch instruction.Op {
		case "nop":
			pc += 1 // do nothing
		case "jmp":
			pc += instruction.Arg
		case "acc":
			accum += instruction.Arg
			pc += 1
		}

	}
	return accum, true
}

func (prog *HandheldProgram) permute(i int) {
	ref := &(*prog)[i]
	switch ref.Op {
	case "jmp":
		ref.Op = "nop"
	case "nop":
		ref.Op = "jmp"
	}
}

// Fix fixes the program (if necessary) returns the fixed instruction and wether id did actually fix
func (prog *HandheldProgram) Fix() (int, bool) {
	if _, ok := prog.Run(); ok {
		return 0, false
	}

	for fixed := 0; fixed < prog.Len(); fixed++ {
		// try
		prog.permute(fixed)
		_, ok := prog.Run()
		if ok {
			return fixed, true
		}
		prog.permute(fixed)
	}
	return 0, false
}
