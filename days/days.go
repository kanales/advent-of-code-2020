package days

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/kanales/advent-of-code-2020/days/day1"
	"github.com/kanales/advent-of-code-2020/days/day13"
	"github.com/kanales/advent-of-code-2020/days/day14"
	"github.com/kanales/advent-of-code-2020/days/day15"
	"github.com/kanales/advent-of-code-2020/days/day2"
	"github.com/kanales/advent-of-code-2020/days/day3"
	"github.com/kanales/advent-of-code-2020/days/day4"
	"github.com/kanales/advent-of-code-2020/days/day5"
	"github.com/kanales/advent-of-code-2020/days/day6"
	"github.com/kanales/advent-of-code-2020/days/day7"
	"github.com/kanales/advent-of-code-2020/days/day8"
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
		expenses := day1.ParseExpenses(input)
		x, y, _ := day1.FindExpenses2(expenses, 2020)
		first := x * y

		x, y, z, _ := day1.FindExpenses3(expenses, 2020)
		second := x * y * z
		return DayResult{Day: 1, First: first, Second: second}
	},
	func(input []byte) DayResult {
		records := day2.ParseRecords(input)
		first := day2.CountCorrectPasswords(records, day2.IsPasswordCorrect1)
		second := day2.CountCorrectPasswords(records, day2.IsPasswordCorrect2)
		return DayResult{Day: 2, First: first, Second: second}
	},
	func(input []byte) DayResult {
		area := day3.ParseSlidingArea(input)
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
		passports := day4.ParsePassports(input)
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
	func(input []byte) DayResult {
		seats := day5.ParsePlaneSeating(input)

		second, err := seats.FindMissingId()
		if err != nil {
			panic(err)
		}
		return DayResult{Day: 5, First: seats.FindMaxId(), Second: second}
	},

	func(input []byte) DayResult {
		groups := day6.ParseCustomsGroups(input)
		first := 0
		for _, group := range groups {
			first += group.CombinedAnswers()
		}
		second := 0
		for _, group := range groups {
			second += group.CommonAnswers()
		}
		return DayResult{Day: 6, First: first, Second: second}
	},

	func(input []byte) DayResult {
		rules := day7.ParseLuggageRules(input)
		first := rules.CountCanContain("shiny gold")
		second := rules.BagsContained("shiny gold")
		return DayResult{Day: 7, First: first, Second: second}
	},

	func(input []byte) DayResult {
		prog := day8.ParseHandheldProgram(input)
		first, _ := prog.Run()
		prog.Fix()
		second, _ := prog.Run()
		return DayResult{Day: 8, First: first, Second: second}
	},
	func(input []byte) DayResult {
		// TODO
		return DayResult{Day: 9, First: 0, Second: 0}
	},
	func(input []byte) DayResult {
		// TODO
		return DayResult{Day: 10, First: 0, Second: 0}
	},
	func(input []byte) DayResult {
		// TODO
		return DayResult{Day: 11, First: 0, Second: 0}
	},
	func(input []byte) DayResult {
		// TODO
		return DayResult{Day: 12, First: 0, Second: 0}
	},
	func(input []byte) DayResult {
		parsed := day13.ParseInput(input)
		sched := parsed.Schedule
		closestID, wait := sched.FindClosestPast(parsed.Timestamp)
		first := closestID * wait

		second := sched.FindTimestampOrder()
		return DayResult{Day: 13, First: int(first), Second: int(second)}
	},
	func(input []byte) DayResult {
		program := day14.ParseProgram(input)
		first := uint64(0)
		for _, b := range program.Run(nil) {
			first += b
		}

		second := uint64(0)
		for _, b := range program.RunV2(nil) {
			second += b
		}
		return DayResult{Day: 14, First: int(first), Second: int(second)}
	},
	func(input []byte) DayResult {
		starting := day15.ParseInput(input)
		first := day15.MemoryGame(starting, 2020)
		second := day15.MemoryGame(starting, 30000000)

		return DayResult{Day: 15, First: int(first), Second: int(second)}
	},
}

// FetchInput outputs the input for that day, or downloads it if necessary
func FetchInput(client *http.Client, year int, day int) ([]byte, error) {
	filename := path.Join(".", "cache", fmt.Sprintf("input_%d.txt", day))

	if fileExists(filename) {
		return ioutil.ReadFile(filename)
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("%v when fetching day %v", resp.StatusCode, day))
	}
	defer resp.Body.Close()

	file, err := os.Create(filename)
	defer file.Close()
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
