package day6

import (
	"bytes"

	"github.com/kanales/advent-of-code-2020/util"
)

type CustomsGroup map[byte]int

func ParseCustomsGroup(input []byte) CustomsGroup {
	out := make(map[byte]int)
	nlines := bytes.Count(input, util.NL) + 1
	for _, c := range input {
		if 'a' <= c && c <= 'z' {
			out[c] += 1
		}
	}

	for k, v := range out {
		if v < nlines {
			out[k] = 0
		} else {
			out[k] = 1
		}
	}

	return out
}

func ParseCustomsGroups(input []byte) []CustomsGroup {
	input = bytes.TrimRight(input, "\n")
	groupBytes := bytes.Split(input, []byte{'\n', '\n'})
	groups := make([]CustomsGroup, len(groupBytes))
	for i, line := range groupBytes {
		groups[i] = ParseCustomsGroup(line)
	}

	return groups
}

func (group *CustomsGroup) CombinedAnswers() int {
	return len(*group)
}

func (group *CustomsGroup) CommonAnswers() int {
	count := 0
	for _, v := range *group {
		count += v
	}
	return count
}
