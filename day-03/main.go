package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	vals := []string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}

	count := countPriorities(vals)
	fmt.Println("count", count)

	count = countPriorities2(vals)
	fmt.Println("count", count)

	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	vals = nil

	filescanner := bufio.NewScanner(readFile)
	filescanner.Split(bufio.ScanLines)

	for filescanner.Scan() {
		val := filescanner.Text()
		vals = append(vals, val)
	}

	count = countPriorities(vals)
	fmt.Println("count", count)

	count = countPriorities2(vals)
	fmt.Println("count", count)

}

func countPriorities(vals []string) int {
	var (
		total int
	)
	for _, val := range vals {
		n := len(val)
		s1 := val[:n/2]
		s2 := val[n/2:]
		kvp1 := make(map[byte]int)
		kvp2 := make(map[byte]int)

		for i := 0; i < len(s1); i++ {
			kvp1[s1[i]]++
		}

		for i := 0; i < len(s2); i++ {
			kvp2[s2[i]]++
		}

		var (
			common byte
			v      byte
		)
		for k := range kvp1 {
			if _, ok := kvp2[k]; ok {
				common = k
				break
			}
		}

		if common >= 'a' && common <= 'z' {
			v = common - 'a' + 1
		} else {
			v = common - 'A' + 27
		}

		vi := int(v)
		total += vi
	}

	return total
}

func countPriorities2(vals []string) int {
	var (
		total        int
		totalSegment = 3
		kvps         = make([]map[byte]int, totalSegment)
		segments     [][]string
		se           = 0
	)
	for i := 0; i < len(vals); i++ {
		if (i+1)%totalSegment == 0 || i == len(vals)-1 {
			segments = append(segments, vals[se:i+1])
			se = i + 1
		}
	}

	for i := 0; i < len(segments); i++ {
		for j := 0; j < len(segments[i]); j++ {
			val := segments[i][j]
			kvps[j] = make(map[byte]int)
			for k := 0; k < len(val); k++ {
				kvps[j][val[k]]++
			}
		}

		var (
			common  byte
			commons []byte
			v       byte
		)

		for k := range kvps[0] {
			if _, ok := kvps[1][k]; ok {
				commons = append(commons, k)
			}
		}

		for i := 0; i < len(commons); i++ {
			c := commons[i]
			if _, ok := kvps[2][c]; ok {
				common = c
				break
			}
		}

		if common >= 'a' && common <= 'z' {
			v = common - 'a' + 1
		} else {
			v = common - 'A' + 27
		}

		vi := int(v)
		total += vi
	}

	return total
}
