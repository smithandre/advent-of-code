package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := readFile("day3/input.txt")
	fmt.Println(input)
}

func readFile(filename string) []string {
	file, err := os.Open(os.Getenv("GOPATH") + "/src/adventofcode/2020/" + filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}