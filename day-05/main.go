package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	input, err := readInput("input-test.txt")
	if err != nil {
		log.Fatal(err)
	}
	topCrates := findTopCrates(*input)
	fmt.Println(topCrates)

	input, err = readInput("input-test.txt")
	if err != nil {
		log.Fatal(err)
	}
	topCrates = findTopCrates2(*input)
	fmt.Println(topCrates)

	input, err = readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	topCrates = findTopCrates(*input)
	fmt.Println(topCrates)

	input, err = readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	topCrates = findTopCrates2(*input)
	fmt.Println(topCrates)
}

func findTopCrates(input Input) string {
	var crate string
	for _, command := range input.Commands {
		n := command.N
		from := command.From - 1
		to := command.To - 1

		for i := 0; i < n; i++ {
			val := input.Stacks[from][0]
			input.Stacks[from] = input.Stacks[from][1:]
			input.Stacks[to] = append([]byte{val}, input.Stacks[to]...)
		}
	}

	for _, stack := range input.Stacks {
		if len(stack) > 0 {
			crate += string(stack[0])
		}
	}

	return crate
}

func findTopCrates2(input Input) string {
	var crate string
	for _, command := range input.Commands {
		n := command.N
		from := command.From - 1
		to := command.To - 1

		var batch Stack

		for i := 0; i < n; i++ {

			val := input.Stacks[from][0]
			input.Stacks[from] = input.Stacks[from][1:]
			batch = append(batch, val)
		}

		for i := n - 1; i >= 0; i-- {

			input.Stacks[to] = append([]byte{batch[i]}, input.Stacks[to]...)
		}
	}

	for _, stack := range input.Stacks {
		if len(stack) > 0 {
			crate += string(stack[0])
		}
	}

	return crate
}

func readInput(filename string) (*Input, error) {
	var (
		input Input
		data  []string
	)
	readFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	filescanner := bufio.NewScanner(readFile)
	filescanner.Split(bufio.ScanLines)

	for filescanner.Scan() {
		val := filescanner.Text()
		if val == "" {
			input.Stacks = createStack(data)
			data = nil
			continue

		}
		data = append(data, val)
	}
	var commands []Command
	for _, d := range data {
		commands = append(commands, parseCommand(d))
	}
	input.Commands = commands
	return &input, err
}

func createStack(data []string) []Stack {
	var (
		stacks  []Stack
		n       = len(data) - 1
		indexes []int
	)
	for i := n; i >= 0; i-- {
		if i == n {
			val := data[i]
			for j := 0; j < len(val); j++ {
				if unicode.IsDigit(rune(val[j])) {
					indexes = append(indexes, j)
				}
			}
			stacks = make([]Stack, len(indexes))
		} else {
			val := data[i]
			l := 0
			for j := 0; j < len(val); j++ {
				if l == len(indexes) {
					break
				}
				if j == indexes[l] {
					if val[j] != ' ' {
						stacks[l] = append([]byte{val[j]}, stacks[l]...)
					}
					l++
				}
			}
		}

	}
	return stacks
}

type Input struct {
	Stacks   []Stack
	Commands []Command
}

type Command struct {
	N    int
	From int
	To   int
}

func parseCommand(s string) Command {
	var (
		command Command
		n       int
		from    int
		to      int
		m       int
		f       int
		t       int
	)
	m = strings.Index(s, "move")
	f = strings.Index(s, "from")
	t = strings.Index(s, "to")

	for i := m; i < f; i++ {
		if unicode.IsDigit(rune(s[i])) {
			n = n*10 + int(s[i]-'0')
		}
	}
	command.N = n

	for i := f; i < t; i++ {
		if unicode.IsDigit(rune(s[i])) {
			from = from*10 + int(s[i]-'0')
		}
	}
	command.From = from

	for i := t; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			to = to*10 + int(s[i]-'0')
		}
	}
	command.To = to

	return command
}

type Stack []byte
