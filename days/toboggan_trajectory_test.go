package days

import (
	"testing"
)

func TestParseSlidingArea(t *testing.T) {
	input := []byte("..##..\n#...#.\n")
	expect := []bool{
		false, false, true, true, false, false,
		true, false, false, false, true, false,
	}

	got := ParseSlidingArea(input)

	if got.Cols != 6 {
		t.Errorf("this.Cols = %v; want %v", got.Cols, 6)
	}

	if got.Rows != 2 {
		t.Errorf("this.Cols = %v; want %v", got.Rows, 2)
	}

	for i, tree := range expect {
		x := i % 6
		y := i / 6

		if got.TreeAt(x, y) != tree {
			t.Errorf("this.TreeAt(%v, %v) = %v; want %v", x, y, got.TreeAt(x, y), tree)
		}
	}
}

func TestCastRay(t *testing.T) {
	input := []byte("..##.......\n#...#...#..\n.#....#..#.\n..#.#...#.#\n.#...##..#.\n..#.##.....\n.#.#.#....#\n.#........#\n#.##...#...\n#...##....#\n.#..#...#.#")

	area := ParseSlidingArea(input)

	tests := [](struct{ x, y, expect int }){
		{x: 1, y: 1, expect: 2},
		{x: 3, y: 1, expect: 7},
		{x: 5, y: 1, expect: 3},
		{x: 7, y: 1, expect: 4},
		{x: 1, y: 2, expect: 2},
	}

	for _, test := range tests {
		got := area.CastRay(test.x, test.y)
		if got != test.expect {
			t.Errorf("this.CastRay(%v, %v) = %v; want %v", test.x, test.y, got, test.expect)
		}
	}

}
