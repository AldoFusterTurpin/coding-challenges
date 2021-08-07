// ALDO FUSTER TURPIN

package main

import "testing"

func EqualSlices(s1, s2 []int32) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func TestEqualSlices(t *testing.T) {
	tt := map[string]struct {
		s1       []int32
		s2       []int32
		expected bool
	}{
		"Empty equal slices": {
			[]int32{},
			[]int32{},
			true,
		},
		"Size one equal slices": {
			[]int32{100},
			[]int32{100},
			true,
		},
		"Equal long slices": {
			[]int32{100, 100, 50, 40, 40, 20, 10},
			[]int32{100, 100, 50, 40, 40, 20, 10},
			true,
		},
		"Different size slices": {
			[]int32{100, 100, 50, 40, 40, 20, 10},
			[]int32{5, 25, 50, 120},
			false,
		},
		"Different long slices": {
			[]int32{100, 100, 50, 40, 40, 20, 10},
			[]int32{90, 4, 50, 40, 40, 20, 89},
			false,
		},
		"Different slices of size one": {
			[]int32{101},
			[]int32{100},
			false,
		},
		"Slice with one element and other empty": {
			[]int32{},
			[]int32{100},
			false,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := EqualSlices(tc.s1, tc.s2)
			if got != tc.expected {
				t.Errorf("Expected %v but got: %v.", tc.expected, got)
			}
		})
	}
}

func TestClimbingLeaderboard(t *testing.T) {
	tt := map[string]struct {
		ranked   []int32
		player   []int32
		expected []int32
	}{
		"Sample 1": {
			[]int32{100, 100, 50, 40, 40, 20, 10},
			[]int32{5, 25, 50, 120},
			[]int32{6, 4, 2, 1},
		},
		"Sample 2": {
			[]int32{100, 90, 90, 80, 75, 60},
			[]int32{50, 65, 77, 90, 102},
			[]int32{6, 5, 4, 2, 1},
		},
		"Custom sample 1": {
			[]int32{100, 80},
			[]int32{110, 100, 90, 80, 70, 60},
			[]int32{1, 1, 2, 2, 3, 3},
		},
		"Custom sample 2": {
			[]int32{110, 90, 60},
			[]int32{110, 100, 90, 80, 70, 50},
			[]int32{1, 2, 2, 3, 3, 4},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := climbingLeaderboard(tc.ranked, tc.player)
			if !EqualSlices(got, tc.expected) {
				t.Errorf("Expected %v but got: %v.", tc.expected, got)
			}
		})
	}
}
