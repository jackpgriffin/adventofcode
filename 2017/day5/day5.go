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
	file, err := os.Open("day5.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := strings.TrimSuffix(scanner.Text(), " ")
		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, int(i))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//input = []int{0, 3, 0, 1, -3}
	inputP1, inputP2 := make([]int, len(input)), make([]int, len(input))
	copy(inputP1, input)
	copy(inputP2, input)

	steps := 0
	i := 0
	for {
		steps++
		oldI := i
		i += inputP1[i]
		if i >= len(inputP1) {
			break
		}
		inputP1[oldI]++
	}
	fmt.Printf("Part 1: %d\n", steps)

	steps = 0
	i = 0
	for {
		if i >= len(inputP2) {
			break
		}

		v := inputP2[i]
		if v >= 3 {
			inputP2[i] -= 1
		} else {
			inputP2[i] += 1
		}

		i += v
		steps++
	}

	fmt.Printf("Part 2: %d\n", steps)
}
