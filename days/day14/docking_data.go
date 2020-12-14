package day14

import (
	"math"
	"regexp"
	"strconv"
)

type Mask struct {
	// masks have 32 < [36] < 64 bits
	Ones, Zeros uint64
}

type Instruction struct {
	Address uint64
	Value   uint64
}
type subProgram struct {
	Mask         Mask
	Instructions []Instruction
}

type Program []subProgram

func ParseMask(input []byte) Mask {
	var ones, zeros uint64

	for _, b := range input {
		zeros <<= 1
		ones <<= 1
		switch b {
		case '1':
			ones++
		case '0':
			zeros++
		}
	}
	return Mask{Ones: ones, Zeros: zeros}
}

func (mask *Mask) Apply(value uint64) uint64 {
	return (value &^ mask.Zeros) | mask.Ones
}

func (mask *Mask) Float(address uint64) []uint64 {
	address = address | mask.Ones

	// indicates the changing bits
	changing := ((1 << 36) - 1) ^ (mask.Ones | mask.Zeros)

	floating := []uint64{address}

	bytecount := math.Ilogb(float64(changing))
	// length of the changing bytes
	for significant := uint64(1 << bytecount); significant > 0; significant >>= 1 {
		n := len(floating)
		for i := 0; i < n; i++ {
			if s := changing & significant; s != 0 {
				floating = append(floating, floating[i]^s)
			}
		}
	}

	return floating
}

var lineRe *regexp.Regexp

func init() {
	lineRe = regexp.MustCompile(`(?m)^mask = ([X10]{36})$|^mem\[(\d+)\] = (\d+)$`)
}

func ParseProgram(input []byte) Program {
	matches := lineRe.FindAllSubmatch(input, -1)

	// parse masks

	maxAddress := 0
	subprogs := make([]subProgram, 0)
	for i := 0; i < len(matches); {
		match := matches[i]

		mask := ParseMask(match[1])
		instructions := make([]Instruction, 0)
		for i++; i < len(matches); i++ {
			match := matches[i]
			if match[1] != nil {
				break
			}
			addr, _ := strconv.Atoi(string(match[2]))
			if addr > maxAddress {
				maxAddress = addr
			}
			value, _ := strconv.Atoi(string(match[3]))
			instructions = append(instructions, Instruction{Address: uint64(addr), Value: uint64(value)})
		}
		subprogs = append(subprogs, subProgram{
			Mask: mask, Instructions: instructions,
		})
	}

	return subprogs
}

func (prog *subProgram) Run(memory map[uint64]uint64) map[uint64]uint64 {
	mask := prog.Mask
	for _, i := range prog.Instructions {
		memory[i.Address] = mask.Apply(i.Value)
	}
	return memory
}

func (prog *Program) Run(memory map[uint64]uint64) map[uint64]uint64 {
	if memory == nil {
		memory = make(map[uint64]uint64)
	}
	for _, p := range *prog {
		p.Run(memory)
	}
	return memory
}

func (prog *subProgram) RunV2(memory map[uint64]uint64) map[uint64]uint64 {
	mask := prog.Mask
	for _, i := range prog.Instructions {
		for _, addr := range mask.Float(i.Address) {
			memory[addr] = i.Value
		}
	}
	return memory
}

func (prog *Program) RunV2(memory map[uint64]uint64) map[uint64]uint64 {
	if memory == nil {
		memory = make(map[uint64]uint64)
	}
	for _, p := range *prog {
		p.RunV2(memory)
	}
	return memory
}
