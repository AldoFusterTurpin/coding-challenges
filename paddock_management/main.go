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
		fmt.Println(err.Error())
	}

	var paddocks []paddock.Paddock
	if err := json.Unmarshal(paddocksJSON, &paddocks); err != nil {
		fmt.Println(err.Error())
	}

	return sps.SolveProblem(paddockTypes, paddocks)
}

func main() {
	paddockTypesJSON, err := ioutil.ReadFile(paddockTypesFile)
	if err != nil {
		fmt.Println("Error", err)
	}

	paddocksJSON, err := ioutil.ReadFile(paddocksFile)
	if err != nil {
		fmt.Println("Error", err)
	}

	result := SolveProblem([]byte(paddockTypesJSON), []byte(paddocksJSON))

	resultJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Printf("Result is:%s\n", string(resultJSON))
}
