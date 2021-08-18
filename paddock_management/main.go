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

func GetResult(paddockTypes []PaddockType, paddocks []Paddock) ResultType {
	return ResultType{}
}

func main() {
	fmt.Println("Main works")
}
