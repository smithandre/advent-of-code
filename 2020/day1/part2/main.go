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

func findValues(lines []int) [3]int {
	var a [3]int
	for i := 0; i < len(lines) - 2; i +=1 {
		for j := i+1; j < len(lines) - 1; j+=1 {
			for k := j+1; k < len(lines); k +=1 {
				if(lines[i] + lines[j] + lines[k] == 2020) {
					fmt.Println("Triplets are ", lines[i], ", ",lines[j], ", ", lines[k])
					a[0] = lines[i]
					a[1] = lines[j]
					a[2] = lines[k]
				}
			}
		}
	}

	return a
}

func calcAnswer(result [3]int) int {
	return result[0] * result[1] * result[2]
}