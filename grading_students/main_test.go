package main

import (
	"testing"
)

func compare(s1 []int32, s2 []int32) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func TestCompareShouldBeTrue(t *testing.T) {
	s1 := []int32{73, 67, 38, 33}
	s2 := []int32{73, 67, 38, 33}

	if !compare(s1, s2) {
		t.Errorf("Compare should return true")
	}
}

func TestCompareShouldBeFalse(t *testing.T) {
	s1 := []int32{73, 67, 0, 33}
	s2 := []int32{733, 627, 38, 33}

	if compare(s1, s2) {
		t.Errorf("Compare should return true")
	}
}

func TestCompareShouldBeFalse2(t *testing.T) {
	s1 := []int32{}
	s2 := []int32{733, 627, 38, 33}

	if compare(s1, s2) {
		t.Errorf("Compare should return true")
	}
}

func TestCompareShouldBeFalse3(t *testing.T) {
	s1 := []int32{}
	s2 := []int32{733, 627, 38, 33}

	if compare(s1, s2) {
		t.Errorf("Compare should return true")
	}
}

func TestGradingStudents(t *testing.T) {
	grades := []int32{73, 67, 38, 33}
	result := gradingStudents(grades)

	expectedGrades := []int32{75, 67, 40, 33}
	if !compare(result, expectedGrades) {
		t.Errorf("Expected %d but got %d", expectedGrades, result)
	}
}

func TestGradingStudentsV2(t *testing.T) {
	grades := []int32{73, 67, 38, 33}
	result := gradingStudents(grades)

	expectedGrades := []int32{75, 67, 40, 33}
	if !compare(result, expectedGrades) {
		t.Errorf("Expected %d but got %d", expectedGrades, result)
	}
}
