package day1

// Day 1 - Report Repair

import (
	"bytes"
	"errors"
	"strconv"
)

// Expenses as a slice of integers
type Expenses []int

// ParseExpenses parses intInput from string
func ParseExpenses(data []byte) Expenses {
	lines := bytes.Split(data, []byte{'\n'})
	lines = lines[:len(lines)-1]
	output := make([]int, len(lines))
	for i, line := range lines {
		value, err := strconv.Atoi(string(line))
		if err != nil {
			panic(err)
		}
		output[i] = value
	}
	return output
}

func FindExpenses2(input Expenses, addto int) (int, int, error) {
	for idx1, elem1 := range input {
		for idx2, elem2 := range input {
			if idx1 != idx2 && elem1+elem2 == addto {
				return elem1, elem2, nil
			}
		}
	}
	return 0, 0, errors.New("Could not find expenses")
}

func FindExpenses3(input Expenses, addto int) (int, int, int, error) {
	for idx1, elem1 := range input {
		for idx2, elem2 := range input {
			for idx3, elem3 := range input {
				if (idx1 != idx2 && idx1 != idx3 && idx2 != idx3) && elem1+elem2+elem3 == addto {
					return elem1, elem2, elem3, nil
				}
			}
		}
	}
	return 0, 0, 0, errors.New("Could not find expenses")
}
