//ALDO FUSTER TURPIN
//problem https://www.hackerrank.com/challenges/divisible-sum-pairs/problem
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)


//aux function I have created
func isDivisible(x1 int32, x2 int32, k int32) bool {
    return ((x1+x2) % k) == int32(0)
}

/*
n: the integer length of array 
ar: an array of integers
k: the integer to divide the pair sum by
*/
// Complete the divisibleSumPairs function below.
func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
    divisibleCount := int32(0)
    for i := int32(0); i < n; i++ {
        for j := i + 1; j < n; j++ {
            if isDivisible(ar[i], ar[j], k) {
                divisibleCount++
            }
        }
    }
    return divisibleCount
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nk := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nk[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    kTemp, err := strconv.ParseInt(nk[1], 10, 64)
    checkError(err)
    k := int32(kTemp)

    arTemp := strings.Split(readLine(reader), " ")

    var ar []int32

    for i := 0; i < int(n); i++ {
        arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
        checkError(err)
        arItem := int32(arItemTemp)
        ar = append(ar, arItem)
    }

    result := divisibleSumPairs(n, k, ar)

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
