package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("day4.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines, dupes, anagrams := 0, 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines++
		words := strings.Split(scanner.Text(), " ")

		// part 1
		w := map[string]bool{}
		for _, word := range words {
			if _, exists := w[word]; exists {
				dupes++
				break
			}
			w[word] = true
		}

		// part 2
		x := map[string]bool{}
		for _, word := range words {
			var r ToRune
			for _, l := range word {
				r = append(r, l)
			}
			sort.Sort(r)
			if _, exists := x[string(r)]; exists {
				anagrams++
				break
			}
			x[string(r)] = true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", lines-dupes)
	fmt.Printf("Part 2: %d\n", lines-anagrams)
}

type ToRune []rune

func (r ToRune) Len() int           { return len(r) }
func (r ToRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ToRune) Less(i, j int) bool { return r[i] < r[j] }
