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

func SolveProblem(paddockTypesJSON, paddocksJSON []byte) string {
	var paddockTypes []paddock.PaddockType
	if err := json.Unmarshal(paddockTypesJSON, &paddockTypes); err != nil {
		panic(err)
	}

	var paddocks []paddock.Paddock
	if err := json.Unmarshal(paddocksJSON, &paddocks); err != nil {
		panic(err)
	}

	sps := paddock.SimplePaddockSolver{}
	result := sps.SolveProblem(paddockTypes, paddocks)

	//resultJSON, err := json.MarshalIndent(result, "", "  ")
	resultJSON, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	return string(resultJSON)
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

	fmt.Printf("Result is: %s\n", result)
}
