package main

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
