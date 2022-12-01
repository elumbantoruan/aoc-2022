package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
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

	topK := topKCalories(totalCalories, 3)

	for _, tk := range topK {
		totalTopThree += tk
	}

	fmt.Println("total top three", totalTopThree)

}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func topKCalories(calories []int, k int) []int {
	h := new(intHeap)
	for _, cal := range calories {
		heap.Push(h, cal)
	}

	result := make([]int, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, heap.Pop(h).(int))
	}
	return result
}

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	tailIndex := h.Len() - 1
	tail := (*h)[tailIndex]
	*h = (*h)[:tailIndex]
	return tail
}
