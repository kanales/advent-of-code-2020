package day16

import "testing"

func TestCountValid(t *testing.T) {
	note := ParseNote([]byte("class: 1-3 or 5-7\nrow: 6-11 or 33-44\nseat: 13-40 or 45-50\nyour ticket:\n7,1,14\nnearby tickets:\n7,3,47\n40,4,50\n55,2,20\n38,6,12\n"))
	expect := 71

	got := note.Validate()
	if got != expect {
		t.Errorf("note.Validate() = %d; want %d", got, expect)
	}
}

func TestOrderFields(t *testing.T) {
	note := ParseNote([]byte("class: 0-1 or 4-19\nrow: 0-5 or 8-19\nseat: 0-13 or 16-19\nyour ticket:\n11,12,13\nnearby tickets:\n3,9,18\n15,1,5\n5,14,9\n"))
	expect := []string{"row", "class", "seat"}
	got := note.OrderedFields()

	for i, f := range expect {
		if got[i].Name != f {
			t.Errorf("got[%d] = %q; want %q", i, got[i].Name, f)
		}
	}
}
