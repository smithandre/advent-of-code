package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main()  {
	file, err := os.Open("./2020/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var lines []int

	for scanner.Scan() {
		var v, _ = strconv.Atoi(scanner.Text())
		lines = append(lines, v)
	}
	var results = findValues(lines)
	fmt.Println(calcAnswer(results))
}

func findValues(lines []int) [2]int {
	var a [2]int
	for i,value := range lines {
		var diff = 2020 - value
		for _,diffVal := range lines[i:] {
			if diffVal == diff {
				a[0] = value
				a[1] = diffVal
				break
			}
		}
	}

	return a
}

func calcAnswer(result [2]int) int {
	return result[0] * result[1]
}