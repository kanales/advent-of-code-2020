package days

import (
	"reflect"
	"testing"
)

/*BFFFBBFRRR: row 70, column 7, seat ID 567.
FFFBBBFRRR: row 14, column 7, seat ID 119.
BBFFBBFRLL: row 102, column 4, seat ID 820.
*/
func TestParsePlaneSeat(t *testing.T) {
	examples := [][]byte{
		[]byte("BFFFBBFRRR"),
		[]byte("FFFBBBFRRR"),
		[]byte("BBFFBBFRLL"),
	}
	expects := []PlaneSeat{
		{Row: 70, Col: 7, Id: 567},
		{Row: 14, Col: 7, Id: 119},
		{Row: 102, Col: 4, Id: 820},
	}

	for i, example := range examples {
		expect := expects[i]
		got := parsePlaneSeat(example)
		if !reflect.DeepEqual(expect, got) {
			t.Errorf("ParsePlaneSeat(%v) = %v; want %v", example, got, expect)
		}
	}
}
