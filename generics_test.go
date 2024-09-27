package main

import (
	"fmt"
	"testing"
)

func Clone[S ~[]E, E any](s S) S {
	return append(s[:0:0], s...)
}

func TestGenerics(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := Clone(a)

	sa, sb := fmt.Sprint(a), fmt.Sprint(b)

	if sa != sb {
		t.Error("must be equal")
	} else {
		t.Logf("a = %s :: b = %s\n", sa, sb)
	}
}
