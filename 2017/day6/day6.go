package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := `11	11	13	7	0	15	5	5	4	4	1	1	7	1	15	11`
	split := regexp.MustCompile(`\s+`)
	b := split.Split(input, -1)
	banks := []int{}
	for _, v := range b {
		i, _ := strconv.Atoi(v)
		banks = append(banks, i)
	}

	configs := map[string]int{}
	steps := 0
	configs[toString(banks)] = steps

	for {
		index := 0
		//fmt.Printf("i: %d - val %d %+v\n", index, banks[index], banks)
		for i, val := range banks {
			if val > banks[index] {
				index = i
			}
		}

		//	fmt.Printf("i: %d - val %d %+v\n", index, banks[index], banks)

		steps++
		redist := banks[index]
		banks[index] = 0
		for i := 0; i < redist; i++ {
			index++
			if index == len(banks) {
				index = 0
			}
			banks[index]++
		}

		newConfig := toString(banks)
		if s, exists := configs[newConfig]; exists {
			fmt.Printf("part 1 steps %d\n", steps)
			fmt.Printf("part 2 cycles %d\n", steps-s)

			break
		}

		configs[newConfig] = steps
	}

}

func toString(in []int) string {
	text := []string{}
	for i := range in {
		n := strconv.Itoa(in[i])
		text = append(text, n)
	}
	return strings.Join(text, "")
}
