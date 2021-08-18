package main

import (
	"reflect"
	"testing"
)

func TestPaddockManager(t *testing.T) {
	tt := map[string]struct {
		paddockTypes []PaddockType
		paddocks     []Paddock
		expected     []ResultType
	}{
		"simple_input": {
			paddockTypes: []PaddockType{
				PaddockType{ID: 1, Name: "PALTOS"},
				PaddockType{ID: 2, Name: "AVELLANOS"},
				PaddockType{ID: 3, Name: "CEREZAS"},
				PaddockType{ID: 4, Name: "NOGALES"},
			},
			paddocks: []Paddock{
				Paddock{PaddockManagerID: 6, FarmID: 1, PaddockTypeID: 1, HarvestYear: 2019, Area: 1200},
				Paddock{PaddockManagerID: 1, FarmID: 3, PaddockTypeID: 4, HarvestYear: 2019, Area: 500},
				Paddock{PaddockManagerID: 5, FarmID: 3, PaddockTypeID: 2, HarvestYear: 2020, Area: 20000},
				Paddock{PaddockManagerID: 2, FarmID: 2, PaddockTypeID: 3, HarvestYear: 2021, Area: 8401},
				Paddock{PaddockManagerID: 3, FarmID: 1, PaddockTypeID: 1, HarvestYear: 2020, Area: 2877},
				Paddock{PaddockManagerID: 5, FarmID: 2, PaddockTypeID: 2, HarvestYear: 2017, Area: 15902},
				Paddock{PaddockManagerID: 3, FarmID: 3, PaddockTypeID: 2, HarvestYear: 2018, Area: 1736},
				Paddock{PaddockManagerID: 2, FarmID: 3, PaddockTypeID: 3, HarvestYear: 2020, Area: 2965},
				Paddock{PaddockManagerID: 4, FarmID: 3, PaddockTypeID: 4, HarvestYear: 2018, Area: 1651},
				Paddock{PaddockManagerID: 5, FarmID: 1, PaddockTypeID: 1, HarvestYear: 2018, Area: 700},
				Paddock{PaddockManagerID: 1, FarmID: 2, PaddockTypeID: 1, HarvestYear: 2019, Area: 7956},
				Paddock{PaddockManagerID: 5, FarmID: 3, PaddockTypeID: 2, HarvestYear: 2020, Area: 3745},
				Paddock{PaddockManagerID: 6, FarmID: 1, PaddockTypeID: 3, HarvestYear: 2021, Area: 11362},
				Paddock{PaddockManagerID: 2, FarmID: 3, PaddockTypeID: 3, HarvestYear: 2021, Area: 300},
				Paddock{PaddockManagerID: 3, FarmID: 2, PaddockTypeID: 2, HarvestYear: 2020, Area: 19188},
				Paddock{PaddockManagerID: 3, FarmID: 1, PaddockTypeID: 1, HarvestYear: 2019, Area: 17137},
				Paddock{PaddockManagerID: 4, FarmID: 3, PaddockTypeID: 2, HarvestYear: 2020, Area: 100},
				Paddock{PaddockManagerID: 2, FarmID: 1, PaddockTypeID: 3, HarvestYear: 2019, Area: 11845},
				Paddock{PaddockManagerID: 5, FarmID: 2, PaddockTypeID: 1, HarvestYear: 2018, Area: 15969},
				Paddock{PaddockManagerID: 1, FarmID: 3, PaddockTypeID: 1, HarvestYear: 2029, Area: 10420},
				Paddock{PaddockManagerID: 5, FarmID: 2, PaddockTypeID: 3, HarvestYear: 2010, Area: 3200},
				Paddock{PaddockManagerID: 6, FarmID: 1, PaddockTypeID: 2, HarvestYear: 2012, Area: 10587},
				Paddock{PaddockManagerID: 2, FarmID: 2, PaddockTypeID: 2, HarvestYear: 2018, Area: 16750},
			},
			expected: []ResultType{
				ResultType{
					PaddockType: PaddockType{
						ID:   2,
						Name: "AVELLANOS",
					},
					HectaresSum: 88008,
				},
				ResultType{
					PaddockType: PaddockType{
						ID:   1,
						Name: "PALTOS",
					},
					HectaresSum: 56259,
				},
				ResultType{
					PaddockType: PaddockType{
						ID:   3,
						Name: "CEREZAS",
					},
					HectaresSum: 38073,
				},
				ResultType{
					PaddockType: PaddockType{
						ID:   4,
						Name: "NOGALES",
					},
					HectaresSum: 2151,
				},
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := GetResult(tc.paddockTypes, tc.paddocks)
			if !reflect.DeepEqual(tc.expected, got) {
				t.Errorf("Expected %v but got: %v.", tc.expected, got)
			}
		})
	}
}
