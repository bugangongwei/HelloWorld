package design

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

/* non-generic implementation */
// func min(a, b float64) float64 {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

/* generic implementation */
func minGeneric[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

/* scale every element in s */
func scale[E constraints.Integer](s []E, c E) []E {
	r := make([]E, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

/* copy s and return s[i]*c, constrained the same type */
type Point []int32

func (p Point) String() string {
	return "Point{" + strconv.Itoa(int(p[0])) + ", " + strconv.Itoa(int(p[1])) + "}"
}

func scaleV2[S ~[]E, E constraints.Integer](s S, c E) S {
	r := make(S, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}
