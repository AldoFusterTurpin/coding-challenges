// ALDO FUSTER TURPIN

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)
// Each spring, it doubles in height. 
// Each summer, its height increases by 1 meter
/*
Period  Height
0          1 init 
1          2 spring
2          3 summer
3          6 spring
4          7 summer
5          14 spring
*/

/*
 * Complete the 'utopianTree' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER n as parameter.
 */

/*
1 <= t <= 10
0 <= n <= 60
*/

func utopianTree(n int32) int32 {
    // Write your code here
    height := int32(1)
    for i := int32(1); i <= n ; i++ {
        if i % 2 == 0 { 
            // summer
            height++
        } else { 
            // spring
            height *= 2
        }
    }
    return height
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        n := int32(nTemp)

        result := utopianTree(n)

        fmt.Fprintf(writer, "%d\n", result)
    }

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
