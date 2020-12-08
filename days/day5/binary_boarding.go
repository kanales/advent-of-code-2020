package day5

import (
	"bytes"
	"errors"
	"fmt"
	"sort"

	"github.com/kanales/advent-of-code-2020/util"
)

type PlaneSeat struct {
	Row int
	Col int
	Id  int
}

func parsePlaneSeat(input []byte) PlaneSeat {
	seat := PlaneSeat{}
	low, high := 0, 128
	for _, b := range input[:7] {
		dist := high - low
		switch b {
		case 'F':
			high -= dist / 2
		case 'B':
			low += dist / 2
		default:
			panic(fmt.Sprintf("Unexpected %q", b))
		}
	}
	seat.Row = low

	low, high = 0, 8
	for _, b := range input[7:] {
		switch b {
		case 'L':
			high -= (high - low) / 2
		case 'R':
			low += (high - low) / 2
		default:
			panic(fmt.Sprintf("Unexpected %q", b))
		}
	}
	seat.Col = low

	seat.Id = seat.Row*8 + seat.Col
	return seat
}

type PlaneSeating []PlaneSeat

func ParsePlaneSeating(input []byte) PlaneSeating {
	lines := bytes.Split(input[:len(input)-1], util.NL)
	seats := make([]PlaneSeat, len(lines))
	for i, line := range lines {
		seats[i] = parsePlaneSeat(line)
	}

	sort.Slice(seats, func(i, j int) bool {
		return seats[i].Id < seats[j].Id
	})
	return seats
}

func (seating *PlaneSeating) FindMaxId() int {
	return (*seating)[len(*seating)-1].Id
}

func (seating *PlaneSeating) FindMissingId() (int, error) {
	lastRow := (*seating)[len(*seating)-1].Row
	for l := 0; l < len(*seating)-1; l++ {
		left, right := (*seating)[l], (*seating)[l+1]
		if left.Row == 0 || right.Row == lastRow {
			continue
		}

		if right.Id-left.Id > 1 {
			return left.Id + 1, nil
		}
	}
	return 0, errors.New("Could not find seat")
}
