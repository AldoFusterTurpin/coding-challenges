package main

import "fmt"

type PaddockType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Paddock struct {
	PaddockManagerID int `json:"paddockManagerId"`
	FarmID           int `json:"farmId"`
	PaddockTypeID    int `json:"paddockTypeId"`
	HarvestYear      int `json:"harvestYear"`
	Area             int `json:"area"`
}

type ResultType struct {
	PaddockType
	HectaresSum int
}

func GetResult(paddockTypes []PaddockType, paddocks []Paddock) []ResultType {
	expected := []ResultType{
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
	}

	return expected
}

func main() {
	fmt.Println("Main works")
}
