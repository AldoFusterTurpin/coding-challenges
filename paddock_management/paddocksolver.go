package main

import "sort"

func SolveProblem(paddockTypes []PaddockType, paddocks []Paddock) []ResultType {
	// key: id of paddockType
	// value: sum of hectares of all paddocks that have as paddockTypeId the key of this entry of mp
	// i.e: mp[i] returns the sum of hectares of all the paddocks that have "i" as paddockTypeId
	mp := getSumOfHectares(paddocks)

	result := calculateResult(mp, paddockTypes)
	sort.Sort(ByHectare(result))

	return result
}

func getSumOfHectares(paddocks []Paddock) map[int]int {
	mp := make(map[int]int)
	for _, paddock := range paddocks {
		// Don't need to check if entry exists. If the requested key doesn't exist,
		// we get the value type's zero value.
		mp[paddock.PaddockTypeID] += paddock.Area
	}
	return mp
}

func calculateResult(mp map[int]int, paddockTypes []PaddockType) []ResultType {
	result := make([]ResultType, len(paddockTypes))

	for i, paddock := range paddockTypes {
		result[i] = ResultType{
			paddock,
			mp[paddock.ID],
		}
	}
	return result
}
