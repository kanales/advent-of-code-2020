package days

import (
	"reflect"
	"testing"
)

// DAY 1
func TestFindExpenses(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}
	x, y, _ := findExpenses2(input, 2020)
	if x*y != 514579 {
		t.Errorf("FindExpenses(%v) = (%d,%d); want (%d, %d)",
			input, x, y, 1721, 299)
	}

	x, y, z, _ := findExpenses3(input, 2020)
	if x*y*z != 241861950 {
		t.Errorf("FindExpenses(%v) = (%d,%d,%d); want (%d, %d, %d)",
			input, x, y, z, 979, 366, 675)
	}
}

func TestParseExpenses(t *testing.T) {
	input := []byte("1721\n979\n366\n299\n675\n1456\n")
	got := ParseExpenses(input)
	expect := Expenses([]int{1721, 979, 366, 299, 675, 1456})
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("ParseExpenses(%q) = %v; want %v", input, got, expect)
	}
}
