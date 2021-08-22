package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/AldoFusterTurpin/coding-challenges/paddock_management/paddock"
)

const (
	paddockTypesFile = "./input/paddockTypes.json"
	paddocksFile     = "./input/paddocks.json"
)

func SolveProblem(paddockTypesJSON, paddocksJSON []byte) []paddock.ResultType {
	sps := paddock.SimplePaddockSolver{}

	var paddockTypes []paddock.PaddockType
	if err := json.Unmarshal(paddockTypesJSON, &paddockTypes); err != nil {
		panic(err)
	}

	var paddocks []paddock.Paddock
	if err := json.Unmarshal(paddocksJSON, &paddocks); err != nil {
		panic(err)
	}

	return sps.SolveProblem(paddockTypes, paddocks)
}

func main() {
	paddockTypesJSON, err := ioutil.ReadFile(paddockTypesFile)
	if err != nil {
		panic(err)
	}

	paddocksJSON, err := ioutil.ReadFile(paddocksFile)
	if err != nil {
		panic(err)
	}

	result := SolveProblem(paddockTypesJSON, paddocksJSON)

	resultJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result is: %s\n", string(resultJSON))
}
