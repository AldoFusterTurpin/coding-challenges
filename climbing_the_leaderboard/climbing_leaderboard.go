// ALDO FUSTER TURPIN

// ranked is sorted in descending order
// ranked: 100 90  80 70 60
// ranks:   1   2  3  4  5
// i:       0   1  2  3  4
// p: 50
//to return: 6

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readNAndSliceOfSizeN() []int32 {
	var n int32
	fmt.Scanf("%d", &n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	lineStr := scanner.Text()
	mySlice := make([]int32, n)

	lineStrNoSapaces := strings.TrimSpace(lineStr)
	sliceOfStrings := strings.Split(lineStrNoSapaces, " ")

	for i := 0; i < int(n); i++ {
		x, err := strconv.ParseInt(sliceOfStrings[i], 10, 64)
		checkError(err)
		mySlice[i] = int32(x)
	}
	return mySlice
}

func main() {
	ranked := readNAndSliceOfSizeN()
	player := readNAndSliceOfSizeN()
	fmt.Println(ranked)
	fmt.Println(player)
	//result := climbingLeaderboard(ranked, player)
	//fmt.Println(result)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func climbingLeaderboard(ranked, player []int32) []int32 {
	playerRanks := make([]int32, len(player))

	rankedNoDuplicates := removeDuplicates(ranked)
	for i, p := range player {
		playerRanks[i] = findPlace(p, rankedNoDuplicates)
	}
	return playerRanks
}

func removeDuplicates(s []int32) []int32 {
	added := make(map[int32]bool)
	sWithoutDuplicates := make([]int32, 0)

	for _, x := range s {
		if !added[x] {
			sWithoutDuplicates = append(sWithoutDuplicates, x)
			added[x] = true
		}
	}
	return sWithoutDuplicates
}

func findPlace(p int32, ranked []int32) int32 {
	if p >= ranked[0] {
		return 1
	}

	rankedLen := len(ranked)
	if p < ranked[rankedLen-1] {
		return int32(rankedLen) + 1
	}
	if p == ranked[rankedLen-1] {
		return int32(rankedLen)
	}
	return findPlaceBinarySearch(p, ranked)
}

func findPlaceBinarySearch(p int32, rankedNoDuplicates []int32) int32 {
	left_i := 0
	right_i := len(rankedNoDuplicates) - 1
	found := false
	var result int32

	for !found {
		mid := (left_i + right_i) / 2

		if p == rankedNoDuplicates[mid] {
			// +1 because first player has rank 1, not 0 (index)
			result = int32(mid) + 1

			found = true
		} else if p < rankedNoDuplicates[mid] && p > rankedNoDuplicates[mid+1] {

			// +1 because first player has rank 1, not 0 (index)
			// +1 because player's score (p) must be at next adjacent position to the right
			result = int32(mid) + 2
			found = true
		} else if p < rankedNoDuplicates[mid] {
			left_i = mid + 1
		} else if p > rankedNoDuplicates[mid] {
			right_i = mid - 1
		}
	}
	return result
}
