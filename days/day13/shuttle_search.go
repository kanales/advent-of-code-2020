package day13

import (
	"bytes"
	"fmt"
	"math"
	"strconv"

	"github.com/kanales/advent-of-code-2020/util"
	umath "github.com/kanales/advent-of-code-2020/util/math"
)

const MISSING = 0

type BusSchedule []int

type ParseResult struct {
	Timestamp int
	Schedule  BusSchedule
}

func ParseInput(input []byte) ParseResult {
	lines := bytes.Split(input, util.NL)
	timestamp, _ := strconv.Atoi(string(lines[0]))
	times := bytes.Split(lines[1], []byte{','})

	schedule := make(BusSchedule, len(times))
	for i, t := range times {
		t := string(t)
		if t == "x" {
			schedule[i] = MISSING
		} else {
			res, _ := strconv.Atoi(t)
			schedule[i] = res
		}

	}
	return ParseResult{Timestamp: timestamp, Schedule: schedule}
}

func (sched BusSchedule) FindClosestPast(timestamp int) (busID int, wait int) {
	minWait, minID := math.MaxInt64, 0
	for _, id := range sched {
		// Calculate time that we must wait
		// notice that we are finding the
		// smallest 'r' where
		// r = id * k - timestamp (mod id)
		// for some integer k
		if id == MISSING {
			continue
		}
		wait := umath.Mod(-timestamp, id)
		if wait < minWait {
			minWait = wait
			minID = id
		}
	}

	return minID, minWait
}

func (sched BusSchedule) FindTimestampOrder() int64 {
	n := len(sched)
	residuals := make([]int64, n)
	mods := make([]int64, n)

	i := 0
	for j, id := range sched {
		if id != MISSING {
			// Each equation is of the form
			// - timestamp = j (mod id)
			residuals[i] = int64(-j)
			mods[i] = int64(id)
			i++
		}
	}
	mods = mods[:i]
	residuals = residuals[:i]

	// assert all mods are pairwise coprime
	for i, mi := range mods {
		for j := i + 1; j < len(mods); j++ {
			if gcd := umath.GCD(mi, mods[j]); gcd != 1 {
				err := fmt.Sprintf("Can't apply CRT: GCD(%d, %d) = %d", mi, mods[j], gcd)
				panic(err)
			}
		}
	}
	timestamp, _ := umath.BigChineseRemainderTheorem(residuals, mods)

	return timestamp // fix the displacement
}
