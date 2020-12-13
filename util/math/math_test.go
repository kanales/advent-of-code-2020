package math

import "testing"

func TestGCD(t *testing.T) {
	type pair struct {
		A, B, G int64
	}
	cases := []pair{
		{A: 10, B: 6, G: 2},
		{A: 7, B: 5, G: 1},
	}

	for _, c := range cases {
		g := GCD(c.A, c.B)
		if g != c.G {
			t.Errorf("ExtendedEuclid(%d, %d) = %d; want %d", c.A, c.B, g, c.G)
		}
	}
}

func TestExtendedEuclid(t *testing.T) {
	type pair struct {
		A, B, X, Y int
	}
	cases := []pair{
		{A: 10, B: 6, X: -1, Y: 2},
		{A: 7, B: 5, X: -2, Y: 3},
	}

	for _, c := range cases {
		x, y := ExtendedEuclid(c.A, c.B)
		if x != c.X || y != c.Y {
			t.Errorf("ExtendedEuclid(%d, %d) = %d, %d; want (%d, %d)", c.A, c.B, x, y, c.X, c.Y)
		}
	}
}

func TestChineseRemainderTheorem(t *testing.T) {
	as := []int{2, 3, 2}
	ms := []int{3, 5, 7}
	var expectT, expectM int = 23, 105
	gotT, gotM := ChineseRemainderTheorem(as, ms)

	if gotT != expectT || gotM != expectM {
		t.Errorf("chineseRemainderTheorem2(...) = %d, %d; want (%d, %d)",
			gotT, gotM, expectT, expectM)
	}

}

func TestBigChineseRemainderTheorem(t *testing.T) {
	as := []int64{2, 3, 2}
	ms := []int64{3, 5, 7}
	var expectT, expectM int64 = 23, 105
	gotT, gotM := BigChineseRemainderTheorem(as, ms)

	if gotT != expectT || gotM != expectM {
		t.Errorf("chineseRemainderTheorem2(...) = %d, %d; want (%d, %d)",
			gotT, gotM, expectT, expectM)
	}

}
