package days

type dayFunc func([]byte) (int, int)

// DayMap maps "days" to functions
var DayMap []dayFunc = []dayFunc{
	func(input []byte) (int, int) {
		expenses := ParseExpenses(input)
		x, y, _ := findExpenses2(expenses, 2020)
		first := x * y

		x, y, z, _ := findExpenses3(expenses, 2020)
		second := x * y * z
		return first, second
	},

	func(input []byte) (int, int) {
		records := ParseRecords(input)
		first := CountCorrectPasswords(records, IsPasswordCorrect1)
		second := CountCorrectPasswords(records, IsPasswordCorrect2)
		return first, second
	},
}
