package paddock

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

// https://pkg.go.dev/sort#example-package
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
