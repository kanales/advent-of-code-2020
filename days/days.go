package days

import (
	"strconv"
)

type dayFunc func(string) (string, string)

// DayMap maps "days" to functions
var DayMap []dayFunc = []dayFunc{
	func(input string) (string, string) {
		expenses := ParseExpenses(input)
		x, y, _ := findExpenses2(expenses, 2020)
		first := strconv.Itoa(x * y)

		x, y, z, _ := findExpenses3(expenses, 2020)
		second := strconv.Itoa(x * y * z)
		return first, second
	},

	func(input string) (string, string) {
		records := ParsePasswords(input)
		first := CountCorrectPasswords(records, IsPasswordCorrect1)
		second := CountCorrectPasswords(records, IsPasswordCorrect2)
		return strconv.Itoa(first), strconv.Itoa(second)
	},
}
