package paddock

import "sort"

type ResultType struct {
	PaddockType
	HectaresSum int
}

type PaddockSolver interface {
	SolveProblem(paddockTypes []PaddockType, paddocks []Paddock) []ResultType
}

type SimplePaddockSolver struct {
}

func (s SimplePaddockSolver) SolveProblem(paddockTypes []PaddockType, paddocks []Paddock) []ResultType {
	// key: id of paddockType
	// value: sum of hectares of all paddocks that have as paddockTypeId the key of this map entry.
	// i.e: mp[i] returns the sum of hectares of all the paddocks that have "i" as paddockTypeId.
	mp := s.getSumOfHectares(paddocks)

	result := s.calculateResult(mp, paddockTypes)
	sort.Sort(ByHectare(result))

	return result
}

func (s SimplePaddockSolver) getSumOfHectares(paddocks []Paddock) map[int]int {
	mp := make(map[int]int)
	for _, paddock := range paddocks {
		// Don't need to check if entry exists. If the requested key doesn't exist,
		// we get the value type's zero value.
		mp[paddock.PaddockTypeID] += paddock.Area
	}
	return mp
}

func (s SimplePaddockSolver) calculateResult(mp map[int]int, paddockTypes []PaddockType) []ResultType {
	result := make([]ResultType, len(paddockTypes))

	for i, paddock := range paddockTypes {
		result[i] = ResultType{
			paddock,
			mp[paddock.ID],
		}
	}
	return result
}
