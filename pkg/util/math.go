package util

import "math"

func AbsInt(a int) int {
	i := float64(a)
	ab := math.Abs(i)
	return int(ab)
}
