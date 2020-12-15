package day15

import (
	"bytes"
	"strconv"
)

func ParseInput(input []byte) []uint {
	nums := bytes.Split(bytes.TrimRight(input, "\n"), []byte{','})
	out := make([]uint, len(nums))
	for i, num := range nums {
		parsed, _ := strconv.ParseUint(string(num), 10, 0)
		out[i] = uint(parsed)
	}
	return out
}

func MemoryGame(starting []uint, n uint) uint {
	// load memory
	round := uint(1)
	memory := make(map[uint]uint, n)
	for ; round <= uint(len(starting)); round++ {
		v := starting[round-1]
		memory[v] = round
	}

	prev := uint(0)
	for ; round < n; round++ {
		var new uint
		if prevRound, ok := memory[prev]; ok {
			new = round - prevRound
		} else {
			new = 0
		}
		memory[prev] = round
		prev = new
	}

	return prev
}
