package paddock

import (
	"reflect"
	"testing"
)

func TestSolveProblem(t *testing.T) {
	tt := map[string]struct {
		paddockTypes []PaddockType
		paddocks     []Paddock
		expected     []ResultType
	}{
		"simple_input": {
			[]PaddockType{
				{1, "PALTOS"},
				{2, "AVELLANOS"},
				{3, "CEREZAS"},
				{4, "NOGALES"},
			},
			[]Paddock{
				{6, 1, 1, 2019, 1200},
				{1, 3, 4, 2019, 500},
				{5, 3, 2, 2020, 20000},
				{2, 2, 3, 2021, 8401},
				{3, 1, 1, 2020, 2877},
				{5, 2, 2, 2017, 15902},
				{3, 3, 2, 2018, 1736},
				{2, 3, 3, 2020, 2965},
				{4, 3, 4, 2018, 1651},
				{5, 1, 1, 2018, 700},
				{1, 2, 1, 2019, 7956},
				{5, 3, 2, 2020, 3745},
				{6, 1, 3, 2021, 11362},
				{2, 3, 3, 2021, 300},
				{3, 2, 2, 2020, 19188},
				{3, 1, 1, 2019, 17137},
				{4, 3, 2, 2020, 100},
				{2, 1, 3, 2019, 11845},
				{5, 2, 1, 2018, 15969},
				{1, 3, 1, 2029, 10420},
				{5, 2, 3, 2010, 3200},
				{6, 1, 2, 2012, 10587},
				{2, 2, 2, 2018, 16750},
			},
			[]ResultType{
				{
					PaddockType{
						2,
						"AVELLANOS",
					},
					88008,
				},
				{
					PaddockType{
						1,
						"PALTOS",
					},
					56259,
				},
				{
					PaddockType{
						3,
						"CEREZAS",
					},
					38073,
				},
				{
					PaddockType{
						4,
						"NOGALES",
					},
					2151,
				},
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			sps := SimplePaddockSolver{}
			got := sps.SolveProblem(tc.paddockTypes, tc.paddocks)
			if !reflect.DeepEqual(tc.expected, got) {
				t.Errorf("Expected %v but got: %v.", tc.expected, got)
			}
		})
	}
}