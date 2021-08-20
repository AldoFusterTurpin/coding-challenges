package paddock

// https://pkg.go.dev/sort#example-package
type ByHectare []ResultType

func (b ByHectare) Len() int {
	return len(b)
}

func (b ByHectare) Less(i, j int) bool {
	return b[i].HectaresSum > b[j].HectaresSum
}

func (b ByHectare) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
