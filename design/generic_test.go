package design

import "testing"

func TestMinGeneric(t *testing.T) {
	type testCase[T any] struct {
		a, b T
		want T
	}

	/* float64 */
	tests := []testCase[float64]{
		{a: 1.0, b: 2.0, want: 1.0},
		{a: 2.0, b: 1.0, want: 1.0},
		{a: -1.0, b: -2.0, want: -2.0},
	}

	for _, tc := range tests {
		got := minGeneric(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("Float64:minGeneric(%v, %v) = %v; want %v", tc.a, tc.b, got, tc.want)
		}
	}

	/* int */
	testInts := []testCase[int]{
		{a: 1, b: 2, want: 1},
		{a: 10, b: 3, want: 3},
		{a: 0, b: -1, want: -1},
	}

	for _, tc := range testInts {
		got := minGeneric(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("Int:minGeneric(%v, %v) = %v; want %v", tc.a, tc.b, got, tc.want)
		}
	}

	/* string */
	testStrings := []testCase[string]{
		{a: "apple", b: "banana", want: "apple"},
		{a: "grape", b: "apple", want: "apple"},
		{a: "zebra", b: "ant", want: "ant"},
	}

	for _, tc := range testStrings {
		got := minGeneric(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("String:minGeneric(%q, %q) = %q; want %q", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestScale(t *testing.T) {
	type testCase[E any] struct {
		s    []E
		c    E
		want []E
	}

	testInts := []testCase[int]{
		{s: []int{1, 2, 3}, c: 2, want: []int{2, 4, 6}},
		{s: []int{-1, -2, -3}, c: 3, want: []int{-3, -6, -9}},
	}

	for _, tc := range testInts {
		got := scale(tc.s, tc.c)
		if len(got) != len(tc.want) {
			t.Errorf("Scale(%v, %v) = %v; want %v", tc.s, tc.c, got, tc.want)
			continue
		}
		for i := range got {
			if got[i] != tc.want[i] {
				t.Errorf("Scale(%v, %v)[%d] = %v; want %v", tc.s, tc.c, i, got[i], tc.want[i])
			}
		}
	}
}

func TestScaleV2(t *testing.T) {
	type testCase[S ~[]E, E any] struct {
		s       S
		c       E
		want    S
		wantStr string
	}

	testPoints := []testCase[Point, int32]{
		{s: Point{1, 2}, c: 2, want: Point{2, 4}, wantStr: "Point{2, 4}"},
		{s: Point{-1, -2}, c: 3, want: Point{-3, -6}, wantStr: "Point{-3, -6}"},
	}

	for _, tc := range testPoints {
		got := scaleV2(tc.s, tc.c)
		if len(got) != len(tc.want) {
			t.Errorf("ScaleV2(%v, %v) = %v; want %v", tc.s, tc.c, got, tc.want)
			continue
		}
		for i := range got {
			if got[i] != tc.want[i] {
				t.Errorf("ScaleV2(%v, %v)[%d] = %v; want %v", tc.s, tc.c, i, got[i], tc.want[i])
			}
		}
		if got.String() != tc.wantStr {
			t.Errorf("ScaleV2(%v, %v).String() = %q; want %q", tc.s, tc.c, got.String(), tc.wantStr)
		}
	}
}
