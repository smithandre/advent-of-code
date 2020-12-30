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
	input := readFile("day2/input.txt")
	var valids int
	for _,val := range input {
		parts := strings.Fields(val)

		searchChar := parts[1][:1]

		s := strings.Split(parts[0], "-")

		min,_ := strconv.Atoi(s[0])
		max,_ := strconv.Atoi(s[1])
		count := strings.Count(parts[2], searchChar)
		if count < min || count > max {
			fmt.Println("Char "+searchChar+" not found between", min, "and", max, "times in "+parts[2])
		} else {
			fmt.Println("Char "+searchChar+" found between", min, "and", max, "times in "+parts[2]+"("+strconv.Itoa(count)+")")
			valids++
		}
	}
	fmt.Println("Valid Passwords", valids)
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