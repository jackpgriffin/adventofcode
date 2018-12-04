package advent

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Claim struct {
	X       int
	Y       int
	Height  int
	Width   int
	Number  string
	Overlap bool
}

func day3() error {
	i, err := os.Open("day3.input")
	if err != nil {
		return err
	}
	defer i.Close()

	claimMap := map[int]map[int]int{}
	claims := []Claim{}
	scanner := bufio.NewScanner(i)
	for scanner.Scan() {
		claim := strings.Split(scanner.Text(), " ")
		XY := strings.Split(claim[2], ",") // 393,863
		sx := XY[0]
		sy := strings.TrimSuffix(XY[1], ":")

		x, _ := strconv.Atoi(sx)
		y, _ := strconv.Atoi(sy)

		sHeight := strings.Split(claim[3], "x")[0]
		sWidth := strings.Split(claim[3], "x")[1]

		height, _ := strconv.Atoi(sHeight)
		width, _ := strconv.Atoi(sWidth)

		c := Claim{
			X:      x + 1,
			Y:      y + 1,
			Height: height,
			Width:  width,
			Number: strings.TrimPrefix(claim[0], "#"),
		}
		claims = append(claims, c)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	for _, claim := range claims {
		for i := 0; i < claim.Height; i++ {
			for j := 0; j < claim.Width; j++ {
				if _, exist := claimMap[claim.X+i]; !exist {
					claimMap[claim.X+i] = map[int]int{}
				}
				claimMap[claim.X+i][claim.Y+j]++
			}
		}
	}

	moreThanTwo := 0
	for i := 0; i <= 1000; i++ {
		for j := 0; j <= 1000; j++ {
			if claimMap[i][j] >= 2 {
				moreThanTwo++
			}
		}
	}

	fmt.Printf("Day 3 part 1 %d\n", moreThanTwo)

	for _, claim := range claims {
		overlaps := false
		for i := 0; i < claim.Height; i++ {
			for j := 0; j < claim.Width; j++ {
				if claimMap[claim.X+i][claim.Y+j] > 1 {
					overlaps = true
				}
			}
		}
		if !overlaps {
			fmt.Printf("Day 3 part 2 %s\n", claim.Number)
			break
		}
	}
	return nil
}
