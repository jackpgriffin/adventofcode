package advent

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day2() error {
	i, err := os.Open("day2.input")
	if err != nil {
		return err
	}
	defer i.Close()

	threes := 0
	twos := 0
	scanner := bufio.NewScanner(i)

	words := []string{}
	for scanner.Scan() {
		seenLetters := map[string]int{}
		letters := strings.Split(scanner.Text(), "")
		for _, letter := range letters {
			seenLetters[letter]++
		}
		two := false
		three := false
		for _, count := range seenLetters {
			if count == 2 {
				//fmt.Printf("Two %s in %s\n", l, scanner.Text())
				two = true
			}
			if count == 3 {
				//fmt.Printf("Three %s in %s\n", l, scanner.Text())
				three = true
			}
		}
		if two {
			twos++
		}
		if three {
			threes++
		}

		words = append(words, scanner.Text())

	}

	fmt.Printf("Day 1 part1: %d\n", twos*threes)

	if err := scanner.Err(); err != nil {
		return err
	}

	diffByOne := []string{}
	for _, wordX := range words {
		for _, wordY := range words {
			diff := 0
			for i, c := range wordX {
				if string(wordY[i]) != string(c) {
					diff++
				}
			}
			if diff == 1 {
				diffByOne = append(diffByOne, wordX)
			}
		}
	}

	// hopefully only2
	wordX := diffByOne[0]
	wordY := diffByOne[1]

	commonLetters := []string{}
	fmt.Printf("%+v\n", diffByOne)
	for i, c := range wordX {
		if string(wordY[i]) == string(c) {
			commonLetters = append(commonLetters, string(c))
		}
	}

	fmt.Printf("%+v\n", commonLetters)
	fmt.Printf("%s\n", strings.Join(commonLetters, ""))

	return nil
}
