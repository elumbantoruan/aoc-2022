package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	compute1()
}

func compute1() {
	lines, err := readInput("inputtest.txt")
	if err != nil {
		log.Fatal(err)
	}
	treesMap := generateTreeMatrix(lines)
	visibles := countVisibility(treesMap)
	fmt.Println(visibles)

	bestScenic := countBestScenic(treesMap)
	fmt.Println(bestScenic)

	lines, err = readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	treesMap = generateTreeMatrix(lines)
	visibles = countVisibility(treesMap)
	fmt.Println(visibles)

	bestScenic = countBestScenic(treesMap)
	fmt.Println(bestScenic)

}

func countVisibility(matrix [][]int) int {
	var (
		rows         = len(matrix)
		cols         = len(matrix[0])
		visibilities = make([][]bool, rows)
		count        = 0
	)

	// initialize
	for r := 0; r < rows; r++ {
		visibilities[r] = make([]bool, cols)
	}

	// left --> right
	for r := 0; r < rows; r++ {
		tallest := -1
		for c := 0; c < cols; c++ {
			if matrix[r][c] > tallest {
				visibilities[r][c] = true
				tallest = matrix[r][c]
			}
		}
	}

	// right --> left
	for r := 0; r < rows; r++ {
		tallest := -1
		for c := cols - 1; c >= 0; c-- {
			if matrix[r][c] > tallest {
				visibilities[r][c] = true
				tallest = matrix[r][c]
			}
		}
	}

	// top --> bottom
	for r := 0; r < rows; r++ {
		tallest := -1
		for c := 0; c < cols; c++ {
			if matrix[c][r] > tallest {
				visibilities[c][r] = true
				tallest = matrix[c][r]
			}
		}
	}

	// bottom --> top
	for r := rows - 1; r >= 0; r-- {
		tallest := -1
		for c := cols - 1; c >= 0; c-- {
			if matrix[c][r] > tallest {
				visibilities[c][r] = true
				tallest = matrix[c][r]
			}
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if visibilities[r][c] {
				count++
			}
		}
	}

	return count

}

func countBestScenic(matrix [][]int) int {
	var (
		rows       = len(matrix)
		cols       = len(matrix[0])
		bestScenic = 0
	)

	// start r = 1 and c = 1 and upper bound rows -1 and cols -1 because best scenic is counted not for trees in the edges
	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			height := matrix[c][r]
			// left
			left := 0
			for i := r - 1; i >= 0; i-- {
				left++
				if height <= matrix[c][i] {
					break
				}
			}

			// right
			right := 0
			for i := r + 1; i < rows; i++ {
				right++
				if height <= matrix[c][i] {
					break
				}
			}

			// top
			top := 0
			for i := c - 1; i >= 0; i-- {
				top++
				if height <= matrix[i][r] {
					break
				}
			}

			// bottom
			bottom := 0
			for i := c + 1; i < cols; i++ {
				bottom++
				if height <= matrix[i][r] {
					break
				}
			}

			score := left * right * top * bottom
			if score > bestScenic {
				bestScenic = score
			}
		}
	}

	return bestScenic
}

func generateTreeMatrix(lines []string) [][]int {
	var (
		treesMap [][]int
		trees    []int
	)

	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			height := int(line[i] - '0')
			trees = append(trees, height)
		}
		treesMap = append(treesMap, trees)
		trees = nil
	}
	return treesMap
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
