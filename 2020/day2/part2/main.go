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

		pos1,_ := strconv.Atoi(s[0])
		pos2,_ := strconv.Atoi(s[1])

		str := strings.Split(parts[2], "")
		if (str[pos1-1] == searchChar && str[pos2-1] != searchChar) || (str[pos1-1] != searchChar && str[pos2-1] == searchChar) {
			fmt.Println("Char "+searchChar+" found at position", pos1, "or", pos2, " in "+parts[2])
			valids++
		} else {
			fmt.Println("Char "+searchChar+" not found at position", pos1, "or", pos2, " in "+parts[2])
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