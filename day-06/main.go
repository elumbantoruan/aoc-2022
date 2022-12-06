package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	values := []string{
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
		"nppdvjthqldpwncqszvftbrmjlhg",
		"bvwbjplbgvbhsrlpgdmjqwftvncz",
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
	}
	for _, val := range values {
		marker := firstMarker(val, 4)
		fmt.Println(marker)
	}

	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	marker := firstMarker(string(bytes), 4)
	fmt.Println(marker)

	marker = firstMarker(string(bytes), 14)
	fmt.Println(marker)
}

func firstMarker(s string, target int) int {

	var (
		hashSet Hashset
		bytes   = []byte(s)
		i       = 0
		start   = target - 1
	)

	for j := start; j < len(bytes); j++ {
		list := bytes[i : j+1]
		hashSet = NewHashSet()
		hashSet.Add(list...)
		if hashSet.Count() == target {
			return j + 1
		}
		i++
	}
	return 0

}

type Hashset map[byte]interface{}

func NewHashSet() Hashset {
	hashSet := make(map[byte]interface{})
	return hashSet
}

func (h Hashset) Add(values ...byte) {
	for _, val := range values {
		h[val] = nil
	}
}

func (h Hashset) Values() []byte {
	var bytes []byte
	for k := range h {
		bytes = append(bytes, k)
	}
	return bytes
}

func (h Hashset) Contains(key byte) bool {
	_, ok := h[key]
	return ok
}

func (h Hashset) Count() int {
	return len(h)
}
