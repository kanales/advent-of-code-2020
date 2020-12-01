package days

import (
	"errors"
	"strconv"
)

func findExpenses2(input IntInput, addto int) (int, int, error) {
	for idx1, elem1 := range input {
		for idx2, elem2 := range input {
			if idx1 != idx2 && elem1+elem2 == addto {
				return elem1, elem2, nil
			}
		}
	}
	return 0, 0, errors.New("Could not find expenses")
}

func findExpenses3(input IntInput, addto int) (int, int, int, error) {
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

type dayFunc func(string) (string, string)

// DayMap maps "days" to functions
var DayMap []dayFunc = []dayFunc{
	func(input string) (first string, second string) {
		x, y, _ := findExpenses2(ParseInput(input), 2020)
		first = strconv.Itoa(x * y)

		x, y, z, _ := findExpenses3(ParseInput(input), 2020)
		second = strconv.Itoa(x * y * z)
		return first, second
	},
}
