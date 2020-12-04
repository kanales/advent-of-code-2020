package days

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

// DayResult contains the result for evaluating the problems in a day
type DayResult struct {
	Day    int
	First  int
	Second int
}
type dayFunc func([]byte) DayResult

// DayMap maps "days" to functions
var DayMap []dayFunc = []dayFunc{
	func(input []byte) DayResult {
		expenses := ParseExpenses(input)
		x, y, _ := findExpenses2(expenses, 2020)
		first := x * y

		x, y, z, _ := findExpenses3(expenses, 2020)
		second := x * y * z
		return DayResult{Day: 1, First: first, Second: second}
	},
	func(input []byte) DayResult {
		records := ParseRecords(input)
		first := CountCorrectPasswords(records, IsPasswordCorrect1)
		second := CountCorrectPasswords(records, IsPasswordCorrect2)
		return DayResult{Day: 2, First: first, Second: second}
	},
	func(input []byte) DayResult {
		area := ParseSlidingArea(input)
		first := area.CastRay(3, 1)
		second := 1
		pairs := [](struct{ x, y int }){
			{1, 1},
			{3, 1},
			{5, 1},
			{7, 1},
			{1, 2},
		}
		for _, pair := range pairs {
			second *= area.CastRay(pair.x, pair.y)
		}
		return DayResult{Day: 3, First: first, Second: second}
	},
	func(input []byte) DayResult {
		passports := ParsePassports(input)
		first := 0
		second := 0
		for _, passport := range passports {
			if passport.HasRequiredFields() {
				first++
			}
			if passport.Validate() {
				second++
			}
		}

		return DayResult{Day: 4, First: first, Second: second}
	},
}

// FetchInput outputs the input for that day, or downloads it if necessary
func FetchInput(client *http.Client, year int, day int) ([]byte, error) {
	filename := path.Join(".", "cache", fmt.Sprintf("input_%d.txt", day))

	if fileExists(filename) {
		return ioutil.ReadFile(filename)
	}

	file, err := os.Create(filename)
	defer file.Close()
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	io.Copy(file, resp.Body)
	file.Seek(0, 0)
	return ioutil.ReadAll(file)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
