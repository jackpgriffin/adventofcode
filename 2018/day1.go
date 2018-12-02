package advent

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func day1() error {
	i, err := os.Open("day1.input")
	if err != nil {
		return err
	}
	defer i.Close()

	var frequency int
	frequencies := []int{}
	reach := map[int]int{0: 1}

	scanner := bufio.NewScanner(i)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		frequency += i
		frequencies = append(frequencies, i)
	}

	fmt.Printf("Day 1 part1: %d\n", frequency)

	if err := scanner.Err(); err != nil {
		return err
	}

	frequency = 0
	part2Done := false
	for part2Done == false {
		for _, f := range frequencies {
			frequency += f
			reach[frequency]++
			if reach[frequency] == 2 {
				fmt.Printf("Day 1 part2: %d\n", frequency)
				part2Done = true
				break
			}
		}
	}

	return nil
}
