//ALDO FUSTER TURPIN

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func getDaysOfFebruary(year int32) (daysOfFebruary int32) {
	if julianCalendar := (year >= 1700 && year <= 1917); julianCalendar {
		if leapYear := year%4 == 0; leapYear {
			return 29
		}
		return 28
	}

	if transitionYear := year == 1918; transitionYear {
		// 28 because February in a normal non leap year have 28 days.
		// - 14 because in 1918 after January 31st we have directly the day 14th of February (we have skipped 14 days, skipped == need to substract them from 28)
		// + 1 because we want to count the day 14th, included in the range.

		// It is similar to ask how many numbers are between 5 and 1, you could say 5 - 1 = 4
		// numbers but the correct answer depends if you consider the first and last numbers to be included in the range [5, 1],
		// in that case, you have to do: 5 - 1 + 1 = 5, the correct answer. Same reasoning applies to the problem when doing 28 -14 + 1.

		return 28 - 14 + 1
	}

	// case Gregorian calendar
	if leapYear := (year%400 == 0) || (year%4 == 0 && year%100 != 0); leapYear {
		return 29
	}
	return 28
}

func getDayOfMonth(year int32) (dayOfMonth int32) {
	daysOfJanuary := int32(31)
	daysOfMarch := int32(31)
	daysOfApril := int32(30)
	daysOfMay := int32(31)
	daysOfJune := int32(30)
	daysOfJuly := int32(31)
	daysOfAugust := int32(31)

	daysOfFebruary := getDaysOfFebruary(year)

	dayOfTheProgrammer := int32(256)
	dayOfMonth = dayOfTheProgrammer - (daysOfJanuary + daysOfFebruary + daysOfMarch + daysOfApril + daysOfMay + daysOfJune + daysOfJuly + daysOfAugust)
	return dayOfMonth
}

func dayOfProgrammer(year int32) string {
	dayStr := strconv.Itoa(int(getDayOfMonth(year)))
	monthStr := "09"
	yearStr := strconv.Itoa(int(year))

	return dayStr + "." + monthStr + "." + yearStr
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	year := int32(yearTemp)

	result := dayOfProgrammer(year)

	fmt.Fprintf(writer, "%s\n", result)

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
