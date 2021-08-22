package main

import (
	"reflect"
	"testing"

	"github.com/AldoFusterTurpin/coding-challenges/paddock_management/paddock"
)

func TestSolveProblem(t *testing.T) {
	tt := map[string]struct {
		paddockTypesJSON []byte
		paddocksJSON     []byte
		expected         []paddock.ResultType
	}{
		"simple_input": {
			[]byte(`[
				{ "id": 1, "name": "PALTOS"},
				{ "id": 2, "name": "AVELLANOS" },
				{ "id": 3, "name": "CEREZAS" },
				{ "id": 4, "name": "NOGALES" }
	  		]`),
			[]byte(`[
				{ "paddockManagerId": 6, "farmId": 1, "paddockTypeId": 1, "harvestYear": 2019, "area": 1200 },
				{ "paddockManagerId": 1, "farmId": 3, "paddockTypeId": 4, "harvestYear": 2019, "area": 500 },
				{ "paddockManagerId": 5, "farmId": 3, "paddockTypeId": 2, "harvestYear": 2020, "area": 20000 },
				{ "paddockManagerId": 2, "farmId": 2, "paddockTypeId": 3, "harvestYear": 2021, "area": 8401},
				{ "paddockManagerId": 3, "farmId": 1, "paddockTypeId": 1, "harvestYear": 2020, "area": 2877 },
				{ "paddockManagerId": 5, "farmId": 2, "paddockTypeId": 2, "harvestYear": 2017, "area": 15902 },
				{ "paddockManagerId": 3, "farmId": 3, "paddockTypeId": 2, "harvestYear": 2018, "area": 1736 },
				{ "paddockManagerId": 2, "farmId": 3, "paddockTypeId": 3, "harvestYear": 2020, "area": 2965 },
				{ "paddockManagerId": 4, "farmId": 3, "paddockTypeId": 4, "harvestYear": 2018, "area": 1651 },
				{ "paddockManagerId": 5, "farmId": 1, "paddockTypeId": 1, "harvestYear": 2018, "area": 700 },
				{ "paddockManagerId": 1, "farmId": 2, "paddockTypeId": 1, "harvestYear": 2019, "area": 7956 },
				{ "paddockManagerId": 5, "farmId": 3, "paddockTypeId": 2, "harvestYear": 2020, "area": 3745 },
				{ "paddockManagerId": 6, "farmId": 1, "paddockTypeId": 3, "harvestYear": 2021, "area": 11362 },
				{ "paddockManagerId": 2, "farmId": 3, "paddockTypeId": 3, "harvestYear": 2021, "area": 300 },
				{ "paddockManagerId": 3, "farmId": 2, "paddockTypeId": 2, "harvestYear": 2020, "area": 19188 },
				{ "paddockManagerId": 3, "farmId": 1, "paddockTypeId": 1, "harvestYear": 2019, "area": 17137 },
				{ "paddockManagerId": 4, "farmId": 3, "paddockTypeId": 2, "harvestYear": 2020, "area": 100 },
				{ "paddockManagerId": 2, "farmId": 1, "paddockTypeId": 3, "harvestYear": 2019, "area": 11845 },
				{ "paddockManagerId": 5, "farmId": 2, "paddockTypeId": 1, "harvestYear": 2018, "area": 15969 },
				{ "paddockManagerId": 1, "farmId": 3, "paddockTypeId": 1, "harvestYear": 2029, "area": 10420 },
				{ "paddockManagerId": 5, "farmId": 2, "paddockTypeId": 3, "harvestYear": 2010, "area": 3200 },
				{ "paddockManagerId": 6, "farmId": 1, "paddockTypeId": 2, "harvestYear": 2012, "area": 10587 },
				{ "paddockManagerId": 2, "farmId": 2, "paddockTypeId": 2, "harvestYear": 2018, "area": 16750 }
			]`),
			[]paddock.ResultType{
				paddock.ResultType{
					PaddockType: paddock.PaddockType{
						ID:   2,
						Name: "AVELLANOS",
					},
					HectaresSum: 88008,
				},
				{
					PaddockType: paddock.PaddockType{
						ID:   1,
						Name: "PALTOS",
					},
					HectaresSum: 56259,
				},
				{
					PaddockType: paddock.PaddockType{
						ID:   3,
						Name: "CEREZAS",
					},
					HectaresSum: 38073,
				},
				{
					PaddockType: paddock.PaddockType{
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
			got := SolveProblem(tc.paddockTypesJSON, tc.paddocksJSON)
			if !reflect.DeepEqual(tc.expected, got) {
				t.Errorf("Expected %v but got: %v.", tc.expected, got)
			}
		})
	}
}
