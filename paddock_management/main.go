package main

import (
	"fmt"
	"sort"
)

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

type ByHectare []ResultType

func (bh ByHectare) Len() int {
	return len(bh)
}

func (bh ByHectare) Less(i, j int) bool {
	return bh[i].HectaresSum > bh[j].HectaresSum
}

func (bh ByHectare) Swap(i, j int) {
	bh[i], bh[j] = bh[j], bh[i]
}

func SolveProblem(paddockTypes []PaddockType, paddocks []Paddock) []ResultType {
	mp := getSumOfHectares(paddocks)
	result := calculateResult(mp, paddockTypes)
	sortResult(result)
	return result
}

func getSumOfHectares(paddocks []Paddock) map[int]int {
	// key: id of paddockType
	// value: sum of hectares of all the paddocks that have as paddockTypeId the key of this entry of mp
	// i.e: mp[i] returns the sum of hectares of all the paddocks that have "i" as paddockTypeId
	mp := make(map[int]int)
	for _, p := range paddocks {
		// Don't need to check if entry exists. If the requested key doesn't exist,
		// we get the value type's zero value.
		mp[p.PaddockTypeID] += p.Area
	}
	return mp
}

func calculateResult(mp map[int]int, paddockTypes []PaddockType) []ResultType {
	result := make([]ResultType, len(paddockTypes))

	for i, paddock := range paddockTypes {
		result[i] = ResultType{
			PaddockType: paddock,
			HectaresSum: mp[paddock.ID],
		}
	}
	return result
}

func sortResult(result []ResultType) {
	sort.Sort(ByHectare(result))
}

func main() {
	fmt.Println("Main works")
}
