package days

import (
	"strconv"
	"strings"
)

// IntInput as a slice of integers
type IntInput []int

// ParseInput parses IntInput from string
func ParseInput(data string) IntInput {
	lines := strings.Split(data, "\n")
	lines = lines[:len(lines)-1]
	output := make([]int, len(lines))
	for i, line := range lines {
		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		output[i] = value
	}
	return output
}
