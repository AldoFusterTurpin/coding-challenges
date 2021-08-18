package main

import (
	"reflect"
	"testing"
)

func TestGetResult(t *testing.T) {
	tt := map[string]struct {
		paddockTypes []PaddockType
		paddocks     []Paddock
		expected     []ResultType
	}{
		"simple_input": {
			paddockTypes: []PaddockType{
				{ID: 1, Name: "PALTOS"},
				{ID: 2, Name: "AVELLANOS"},
				{ID: 3, Name: "CEREZAS"},
				{ID: 4, Name: "NOGALES"},
			},
			paddocks: []Paddock{
				{PaddockManagerID: 6, FarmID: 1, PaddockTypeID: 1, HarvestYear: 2019, Area: 1200},
				{PaddockManagerID: 1, FarmID: 3, PaddockTypeID: 4, HarvestYear: 2019, Area: 500},
				{PaddockManagerID: 5, FarmID: 3, PaddockTypeID: 2, HarvestYear: 2020, Area: 20000},
				{PaddockManagerID: 2, FarmID: 2, PaddockTypeID: 3, HarvestYear: 2021, Area: 8401},
				{PaddockManagerID: 3, FarmID: 1, PaddockTypeID: 1, HarvestYear: 2020, Area: 2877},
				{PaddockManagerID: 5, FarmID: 2, PaddockTypeID: 2, HarvestYear: 2017, Area: 15902},
				{PaddockManagerID: 3, FarmID: 3, PaddockTypeID: 2, HarvestYear: 2018, Area: 1736},
				{PaddockManagerID: 2, FarmID: 3, PaddockTypeID: 3, HarvestYear: 2020, Area: 2965},
				{PaddockManagerID: 4, FarmID: 3, PaddockTypeID: 4, HarvestYear: 2018, Area: 1651},
				{PaddockManagerID: 5, FarmID: 1, PaddockTypeID: 1, HarvestYear: 2018, Area: 700},
				{PaddockManagerID: 1, FarmID: 2, PaddockTypeID: 1, HarvestYear: 2019, Area: 7956},
				{PaddockManagerID: 5, FarmID: 3, PaddockTypeID: 2, HarvestYear: 2020, Area: 3745},
				{PaddockManagerID: 6, FarmID: 1, PaddockTypeID: 3, HarvestYear: 2021, Area: 11362},
				{PaddockManagerID: 2, FarmID: 3, PaddockTypeID: 3, HarvestYear: 2021, Area: 300},
				{PaddockManagerID: 3, FarmID: 2, PaddockTypeID: 2, HarvestYear: 2020, Area: 19188},
				{PaddockManagerID: 3, FarmID: 1, PaddockTypeID: 1, HarvestYear: 2019, Area: 17137},
				{PaddockManagerID: 4, FarmID: 3, PaddockTypeID: 2, HarvestYear: 2020, Area: 100},
				{PaddockManagerID: 2, FarmID: 1, PaddockTypeID: 3, HarvestYear: 2019, Area: 11845},
				{PaddockManagerID: 5, FarmID: 2, PaddockTypeID: 1, HarvestYear: 2018, Area: 15969},
				{PaddockManagerID: 1, FarmID: 3, PaddockTypeID: 1, HarvestYear: 2029, Area: 10420},
				{PaddockManagerID: 5, FarmID: 2, PaddockTypeID: 3, HarvestYear: 2010, Area: 3200},
				{PaddockManagerID: 6, FarmID: 1, PaddockTypeID: 2, HarvestYear: 2012, Area: 10587},
				{PaddockManagerID: 2, FarmID: 2, PaddockTypeID: 2, HarvestYear: 2018, Area: 16750},
			},
			expected: []ResultType{
				{
					PaddockType: PaddockType{
						ID:   2,
						Name: "AVELLANOS",
					},
					HectaresSum: 88008,
				},
				{
					PaddockType: PaddockType{
						ID:   1,
						Name: "PALTOS",
					},
					HectaresSum: 56259,
				},
				{
					PaddockType: PaddockType{
						ID:   3,
						Name: "CEREZAS",
					},
					HectaresSum: 38073,
				},
				{
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
