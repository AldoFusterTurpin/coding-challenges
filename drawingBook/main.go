//Aldo Fuster Turpin

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func jumpsToPStartingFromBeginning(n int32, p int32) int32 {
	if p%2 == 0 {
		return p / 2
	}
	return (p - 1) / 2
}

func jumpsToPStartingFromEnd(n int32, p int32) int32 {
	// same page
	if p == n || (p%2 == 0 && n == p+1) {
		return 0
	}

	if n%2 == 0 && p%2 != 0 {
		return (n + 1 - p) / 2
	}

	return (n - p) / 2
}

// int n: the number of pages in the book
// int p: the page number to turn to
// 1 <= n <= 10^5
// 1 <= p <= n
/* If the book is n pages long, and a student wants to turn to page p,
what is the minimum number of pages to turn? They can start at the beginning or the end of the book. */
func pageCount(n int32, p int32) int32 {
	j1 := jumpsToPStartingFromBeginning(n, p)
	j2 := jumpsToPStartingFromEnd(n, p)

	if j1 <= j2 {
		return j1
	}
	return j2
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	pTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	p := int32(pTemp)

	result := pageCount(n, p)

	fmt.Fprintf(writer, "%d\n", result)

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
