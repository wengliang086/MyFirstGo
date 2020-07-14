package main

import (
	"math"
	"testing"
)

func TestAdd2(t *testing.T) {
	tests := []struct {
		a, b, c int
	}{
		{1, 2, 3},
		{0, 2, 2},
		{-1, 1, 0},
		{math.MaxInt64, 1, math.MinInt64},
	}

	for _, test := range tests {
		if actual := Add(test.a, test.b); actual != test.c {
			t.Errorf("error! actual: %d, expacted: %d", actual, test.c)
		}
	}
}
