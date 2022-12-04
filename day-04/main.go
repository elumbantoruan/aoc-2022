package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		data   []string
		merged int
	)

	data = testdata()
	merged = countMerge(data)
	fmt.Println("num of merged", merged)

	merged = countMerge2(data)
	fmt.Println("num of merged2", merged)

	data = getInputData()
	merged = countMerge(data)
	fmt.Println("num of merged", merged)

	merged = countMerge2(data)
	fmt.Println("num of merged2", merged)
}

func countMerge(data []string) int {
	var merged int
	for _, d := range data {
		items := strings.Split(d, ",")
		left, right := items[0], items[1]

		leftItems := strings.Split(left, "-")
		leftStart, _ := strconv.Atoi(leftItems[0])
		leftEnd, _ := strconv.Atoi(leftItems[1])

		rightItems := strings.Split(right, "-")
		rightStart, _ := strconv.Atoi(rightItems[0])
		rightEnd, _ := strconv.Atoi(rightItems[1])

		if leftStart <= rightStart && leftEnd >= rightEnd ||
			leftStart >= rightStart && leftEnd <= rightEnd {
			merged++
		}

	}
	return merged
}

func countMerge2(data []string) int {
	var merged int
	for _, d := range data {
		items := strings.Split(d, ",")
		left, right := items[0], items[1]

		leftItems := strings.Split(left, "-")
		leftStart, _ := strconv.Atoi(leftItems[0])
		leftEnd, _ := strconv.Atoi(leftItems[1])

		rightItems := strings.Split(right, "-")
		rightStart, _ := strconv.Atoi(rightItems[0])
		rightEnd, _ := strconv.Atoi(rightItems[1])

		// swapping
		if rightEnd < leftStart {
			leftStart, rightStart = rightStart, leftStart
			leftEnd, rightEnd = rightEnd, leftEnd
		}

		if leftEnd-rightStart >= 0 {
			merged++
		}

	}
	return merged
}

func getInputData() []string {
	var data []string
	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	filescanner := bufio.NewScanner(readFile)
	filescanner.Split(bufio.ScanLines)

	for filescanner.Scan() {
		val := filescanner.Text()
		data = append(data, val)
	}
	return data
}

func testdata() []string {
	return []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}
}
