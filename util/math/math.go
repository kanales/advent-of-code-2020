package math

import (
	"math/big"
)

// Mod calculates a (mod b)
func Mod(a, b int) int {
	return (a%b + b) % b
}

// ExtendedEuclid finds x,y such that a * x + b * y = GCD(a,b)
func ExtendedEuclid(a, b int) (x, y int) {
	// TODO(kanales): should be optimized
	// Return x, y such that
	//   x * a + y * b = GCD(a, b)
	var oldR, r int = a, b
	var oldS, s int = 1, 0
	var oldT, t int = 0, 1

	for r != 0 {
		quot := oldR / r
		oldR, r = r, oldR-quot*r
		oldS, s = s, oldS-quot*s
		oldT, t = t, oldT-quot*t
	}

	return oldS, oldT
}

// GCD finds the greatest common divisor of `a` and `b`
func GCD(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func BigChineseRemainderTheorem(as, ms []int64) (int64, int64) {
	if len(as) != len(ms) || len(as) < 2 {
		panic("ChineseRemainderTheorem: expected at least 2 equations")
	}

	x := big.NewInt(0)
	y := big.NewInt(0)

	n := big.NewInt(1)
	t := big.NewInt(0)

	for i := 0; i < len(as); i++ {
		a := big.NewInt(as[i])
		m := big.NewInt(ms[i])

		x.ModInverse(n, m)
		y.ModInverse(m, n)

		x.Mul(a, x)
		x.Mul(x, n)

		y.Mul(t, y)
		y.Mul(y, m)

		t.Add(x, y)

		n.Mul(n, m)

		t.Mod(t, n)
		t.Add(t, n)
		t.Mod(t, n)
	}

	return t.Int64(), n.Int64()
}

// ChineseRemainderTheorem finds (T, M) such that
// 	T == as[i] (mod ms[i])
// for all i
func ChineseRemainderTheorem(as, ms []int) (int, int) {
	// fold chinese remainder theorem cases
	if len(as) != len(ms) || len(as) < 2 {
		panic("ChineseRemainderTheorem: expected at least 2 equations")
	}

	t, m := 0, 1

	for i := 0; i < len(as); i++ {
		x, y := ExtendedEuclid(m, ms[i])

		t = t*ms[i]*y + as[i]*m*x

		m = m * ms[i]
		t = Mod(t, m)
	}

	return t, m
}
