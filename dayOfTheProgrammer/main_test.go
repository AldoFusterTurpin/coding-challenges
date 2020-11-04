package main

import (
	"testing"
)

func TestGetDaysOfFebruaryShouldReturn28(t *testing.T) {
	year := int32(2017)
	expected := int32(28)

	if result := getDaysOfFebruary(year); result != expected {
		t.Error("Expected ", expected, ", but got", result)
	}
}

func TestGetDaysOfFebruaryShouldReturn29(t *testing.T) {
	year := int32(2016)
	expected := int32(29)

	if result := getDaysOfFebruary(year); result != expected {
		t.Error("Expected ", expected, ", but got", result)
	}
}

func TestGetDayOfMonthShouldReturn13(t *testing.T) {
	year := int32(2017)
	expected := int32(13)

	if result := getDayOfMonth(year); result != expected {
		t.Error("Expected ", expected, "but got", result)
	}
}

func TestGetDayOfMonthShouldReturn12WhenYearIs2016(t *testing.T) {
	year := int32(2016)
	expected := int32(12)

	if result := getDayOfMonth(year); result != expected {
		t.Error("Expected ", expected, "but got", result)
	}
}

func TestGetDayOfMonthShouldReturn12WhenYearIs1800(t *testing.T) {
	year := int32(1800)
	expected := int32(12)

	if result := getDayOfMonth(year); result != expected {
		t.Error("Expected ", expected, "but got", result)
	}
}

func TestDayOfProgrammerWhenYearIs2017(t *testing.T) {
	year := int32(2017)
	expected := "13.09.2017"

	if result := dayOfProgrammer(year); result != expected {
		t.Error("Expected ", expected, "but got", result)
	}
}

func TestDayOfProgrammerWhenYearIs2016(t *testing.T) {
	year := int32(2016)
	expected := "12.09.2016"

	if result := dayOfProgrammer(year); result != expected {
		t.Error("Expected ", expected, "but got", result)
	}
}

func TestDayOfProgrammerWhenYearIs1800(t *testing.T) {
	year := int32(1800)
	expected := "12.09.1800"

	if result := dayOfProgrammer(year); result != expected {
		t.Error("Expected ", expected, "but got", result)
	}
}

func TestDayOfProgrammerWhenYearIs1918(t *testing.T) {
	year := int32(1918)
	expected := "26.09.1918"

	if result := dayOfProgrammer(year); result != expected {
		t.Error("Expected ", expected, "but got", result)
	}
}
