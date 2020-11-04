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

func didBrianOverChargedAnna(bill []int32, k int32, b int32) (wasOverCharged bool, refundToAnna int32) {
	sum := int32(0)
	lengthOfBill := int32(len(bill))

	for i := int32(0); i < lengthOfBill; i++ {
		if i != k {
			sum += bill[i]
		}
	}
	bActual := sum / 2

	if b == bActual {
		return false, 0
	}

	return true, b - bActual
}

// bill: an array of integers representing the cost of each item ordered
// k: an integer representing the zero-based index of the item Anna doesn't eat
// b: the amount of money that Anna contributed to the bill
func bonAppetit(bill []int32, k int32, b int32) {
	wasOverCharged, amountOvercharged := didBrianOverChargedAnna(bill, k, b)

	if wasOverCharged {
		fmt.Println(amountOvercharged)
	} else {
		fmt.Println("Bon Appetit")
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nk := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	billTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var bill []int32

	for i := 0; i < int(n); i++ {
		billItemTemp, err := strconv.ParseInt(billTemp[i], 10, 64)
		checkError(err)
		billItem := int32(billItemTemp)
		bill = append(bill, billItem)
	}

	bTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	b := int32(bTemp)

	bonAppetit(bill, k, b)
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
