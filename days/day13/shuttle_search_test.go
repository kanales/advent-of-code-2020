package day13

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := []byte("939\n7,13,x,x,59,x,31,19\n")
	expect := ParseResult{Timestamp: 939, Schedule: []int{7, 13, MISSING, MISSING, 59, MISSING, 31, 19}}

	got := ParseInput(input)
	if got.Timestamp != expect.Timestamp {
		t.Errorf("ParseInput(...).Timestamp = %v; want %v", got.Timestamp, expect.Timestamp)
	}

	if !reflect.DeepEqual(got.Schedule, expect.Schedule) {
		t.Errorf("ParseInput(...).Schedule = %v; want %v", got.Schedule, expect.Schedule)
	}
}

func TestFindClosestPast(t *testing.T) {
	input := ParseInput([]byte("939\n7,13,x,x,59,x,31,19\n"))
	expect := 59
	got, _ := input.Schedule.FindClosestPast(input.Timestamp)
	if got != expect {
		t.Errorf("BusSchedule.FindClosestPast(%d) = %d; want %d", input.Timestamp, got, expect)
	}
}

func TestFindTimestampOrder(t *testing.T) {
	type testCase struct {
		Sched  BusSchedule
		Expect int
	}

	cases := []testCase{
		{Sched: []int{7, 13, 0, 0, 59, 0, 31, 19}, Expect: 1068781},
		{Sched: []int{17, 0, 13, 19}, Expect: 3417},
		{Sched: []int{67, 7, 59, 61}, Expect: 754018},
		{Sched: []int{67, 0, 7, 59, 61}, Expect: 779210},
		{Sched: []int{67, 7, 0, 59, 61}, Expect: 1261476},
		{Sched: []int{1789, 37, 47, 1889}, Expect: 1202161486},
	}

	for _, c := range cases {
		got := c.Sched.FindTimestampOrder()
		if got != c.Expect {
			t.Errorf("BusSchedule.FindTimestampOrder() = %d; want %d", got, c.Expect)
		}
	}

}
