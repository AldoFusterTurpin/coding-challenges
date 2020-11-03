package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func getNextMultipleOfFive(n int32) int32 {
	for n%5 != 0 {
		n++
	}
	return n
}

func roundGrade(grade int32) int32 {
	if grade < 38 {
		return grade
	}

	nextMultipleOfFive := getNextMultipleOfFive(grade)
	if nextMultipleOfFive-grade < 3 {
		return nextMultipleOfFive
	}

	return grade
}

func gradingStudents(grades []int32) []int32 {
	roundedGrades := make([]int32, len(grades))

	for i, value := range grades {
		roundedGrades[i] = roundGrade(value)
	}
	return roundedGrades
}

func gradingStudentsV2(grades []int32) []int32 {
	roundedGrades := []int32{}

	for _, value := range grades {
		roundedGrades = append(roundedGrades, roundGrade(value))
	}
	return roundedGrades
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	gradesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var grades []int32

	for i := 0; i < int(gradesCount); i++ {
		gradesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		gradesItem := int32(gradesItemTemp)
		grades = append(grades, gradesItem)
	}

	result := gradingStudents(grades)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
