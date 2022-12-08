package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	data, err := readInput("inputtest.txt")
	if err != nil {
		log.Fatal(err)
	}

	fs := NewFSInfo(data)
	fsum := fs.compute1()
	fmt.Println(fsum)

	fsum2 := fs.compute2()
	fmt.Println(fsum2)

	data, err = readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fs = NewFSInfo(data)
	fsum = fs.compute1()
	fmt.Println(fsum)

	fsum2 = fs.compute2()
	fmt.Println(fsum2)

}

type FSInfo struct {
	Parent      *FSInfo
	FSInfos     []*FSInfo
	Size        int
	Name        string
	IsDirectory bool
}

func (fs FSInfo) SumFileContent() []int {
	var (
		sizes []int
		sum   int
	)
	for _, fi := range fs.FSInfos {
		if fi.IsDirectory {
			fisizes := fi.SumFileContent()
			sum += fisizes[len(fisizes)-1]
			sizes = append(sizes, fisizes...)
		} else {
			sum += fi.Size
		}
	}
	sizes = append(sizes, sum)
	return sizes
}

func (fs FSInfo) compute1() int {
	var (
		target = 100000
		sum    = 0
	)
	sums := fs.SumFileContent()
	sort.Ints(sums)
	for _, size := range sums {
		if size < target {
			sum += size
		}
	}
	return sum
}

func (fs FSInfo) compute2() int {

	sums := fs.SumFileContent()
	sort.Ints(sums)
	shouldBeDeleted := 30000000 - (70000000 - sums[len(sums)-1])
	for _, size := range sums {
		if size > shouldBeDeleted {
			return size
		}
	}
	return 0
}

func NewFSInfo(data []string) FSInfo {
	var (
		root    FSInfo
		current *FSInfo
	)

	for _, line := range data {
		var command string
		if strings.HasPrefix(line, "$") {
			command = strings.TrimPrefix(line, "$ ")
			if strings.HasPrefix(command, "cd") {
				cmd := strings.TrimPrefix(command, "cd ")
				if cmd == "/" {
					root = FSInfo{Name: "/"}
					current = &root
				} else if cmd == ".." {
					current = current.Parent
				} else {
					fs := FSInfo{Name: cmd, Parent: current, IsDirectory: true}
					current.FSInfos = append(current.FSInfos, &fs)
					current = &fs
				}
			}
		} else {
			if strings.HasPrefix(line, "dir") {
				// just print the dir name
				// dir has been created in cd (above
				continue
			}
			var (
				size int
				name string
			)
			_, err := fmt.Sscanf(line, "%d %s", &size, &name)
			if err != nil {
				log.Fatal(err)
			}
			fs := FSInfo{Name: name, Size: size, Parent: current}
			current.FSInfos = append(current.FSInfos, &fs)
		}
	}
	return root
}

func readInput(filename string) ([]string, error) {
	var (
		data []string
	)
	readFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	filescanner := bufio.NewScanner(readFile)
	filescanner.Split(bufio.ScanLines)

	for filescanner.Scan() {
		val := filescanner.Text()

		data = append(data, val)
	}

	return data, nil
}
