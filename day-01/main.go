package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	filescanner := bufio.NewScanner(readFile)
	filescanner.Split(bufio.ScanLines)

	var (
		calorie       int
		totalCalories []int
		totalCalorie  int
		maxCalorie    = math.MinInt
		totalTopThree int
	)

	for filescanner.Scan() {
		val := filescanner.Text()
		if len(val) > 0 {
			calorie, _ = strconv.Atoi(val)
			totalCalorie += calorie
		} else {
			maxCalorie = max(maxCalorie, totalCalorie)
			totalCalories = append(totalCalories, totalCalorie)
			totalCalorie = 0
		}
	}

	fmt.Println("max calorie", maxCalorie)

	sort.Ints(totalCalories)
	n := len(totalCalories)

	totalTopThree += totalCalories[n-1]
	totalTopThree += totalCalories[n-2]
	totalTopThree += totalCalories[n-3]

	fmt.Println("total top three", totalTopThree)

}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
