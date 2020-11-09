//ALDO FUSTER TURPIN

package main

import (
	"testing"
)

// n and p are defined by the problem statement
// int n: the number of pages in the book
// int p: the page number to turn to

func TestJumpsToPStartingFromBeginning(t *testing.T) {
	tests := map[string]struct {
		n    int32
		p    int32
		want int32
	}{
		"n=1;p=1":   {n: 1, p: 1, want: 0},
		"n=45;p=1":  {n: 45, p: 1, want: 0},
		"n=2;p=1":   {n: 2, p: 1, want: 0},
		"n=3;p=1":   {n: 3, p: 1, want: 0},
		"n=4;p=1":   {n: 4, p: 1, want: 0},
		"n=5;p=1":   {n: 5, p: 1, want: 0},
		"n=137;p=1": {n: 137, p: 1, want: 0},
		"n=2;p=2":   {n: 2, p: 2, want: 1},
		"n=3;p=3":   {n: 3, p: 3, want: 1},
		"n=4;p=4":   {n: 4, p: 4, want: 2},
		"n=5;p=5":   {n: 5, p: 5, want: 2},
		"n=6;p=6":   {n: 6, p: 6, want: 3},
		"n=7;p=6":   {n: 7, p: 7, want: 3},
		"n=234;p=6": {n: 234, p: 8, want: 4},
		"n=197;p=6": {n: 197, p: 23, want: 11},
		"n=13;p=4":  {n: 28, p: 13, want: 6},
		"n=19;p=46": {n: 46, p: 19, want: 9},
	}

	// 'range' returns key (the string), value (the struct)
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := jumpsToPStartingFromBeginning(tc.n, tc.p)
			if tc.want != got {
				t.Errorf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}

func TestJumpsToPStartingFromEnd(t *testing.T) {
	tests := map[string]struct {
		n    int32
		p    int32
		want int32
	}{
		"n=1;p=1": {n: 1, p: 1, want: 0},
		"n=2;p=1": {n: 2, p: 1, want: 1},
		"n=3;p=1": {n: 3, p: 1, want: 1},
		"n=4;p=1": {n: 4, p: 1, want: 2},
		"n=5;p=1": {n: 5, p: 1, want: 2},
		"n=2;p=2": {n: 2, p: 2, want: 0},
		"n=3;p=3": {n: 3, p: 3, want: 0},
		"n=4;p=4": {n: 4, p: 4, want: 0},
		"n=5;p=5": {n: 5, p: 5, want: 0},
		"n=6;p=6": {n: 6, p: 6, want: 0},
		"n=7;p=6": {n: 7, p: 7, want: 0},
	}

	// 'range' returns key (the string), value (the struct)
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := jumpsToPStartingFromEnd(tc.n, tc.p)
			if tc.want != got {
				t.Errorf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}

func TestPageCountInput0(t *testing.T) {
	n := int32(5)
	p := int32(3)

	result := pageCount(n, p)

	expected := int32(1)

	if result != expected {
		t.Error("Expected", expected, "but got", result)
	}
}

func TestPageCountInput1(t *testing.T) {
	n := int32(6)
	p := int32(2)

	result := pageCount(n, p)

	expected := int32(1)

	if result != expected {
		t.Error("Expected", expected, "but got", result)
	}
}

func TestPageCountInput2(t *testing.T) {
	n := int32(5)
	p := int32(4)

	result := pageCount(n, p)

	expected := int32(0)

	if result != expected {
		t.Error("Expected", expected, "but got", result)
	}
}
