package advent

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func day5() error {
	i, err := os.Open("day5.input")
	if err != nil {
		return err
	}
	defer i.Close()

	var text string
	scanner := bufio.NewScanner(i)
	for scanner.Scan() {
		text += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	part1 := text
	for {
		newText := removeDups(part1)
		if newText == part1 {
			break
		}
		part1 = newText
	}

	fmt.Printf("Day 5 part 1: %d\n", len(part1))

	part2 := text

	bestPoly := 0
	for i := 'A'; i <= 'Z'; i++ {
		x := removeChar(rune(i), part2)
		for {
			newText := removeDups(x)
			if newText == x {
				break
			}
			x = newText
		}
		if bestPoly == 0 || len(x) < bestPoly {
			bestPoly = len(x)
		}
	}

	fmt.Printf("Day 5 part 2: %d\n", bestPoly)

	return nil
}

func removeChar(char rune, s string) string {
	buf := []rune{}
	for _, c := range []rune(s) {
		if c == char+32 || c == char {
			continue
		}
		buf = append(buf, c)
	}
	return string(buf)
}

func removeDups(s string) string {
	buf := []rune(s)
	i := 0
	for i < len(buf)-1 {
		var c rune
		next := rune(buf[i+1])
		r := rune(buf[i])
		if unicode.IsUpper(r) && unicode.IsLower(next) {
			c = r + 32
		} else if unicode.IsLower(r) && unicode.IsUpper(next) {
			c = r - 32
		}

		if c == next {
			buf = append(buf[:i], buf[i+2:]...)
			i += 2
			continue
		}
		i++
	}
	return string(buf)
}
