package main

import (
	"fmt"
	"math"
)

func main() {
	var input int
	input = 1
	input = 12
	input = 23
	input = 1024
	input = 368078
	res := 0

	for i := 1; i <= input; i += 2 {
		if i*i > input {
			numY := (i - 1) / 2
			innerCorner := (i - 2) * (i - 2)
			layerPos := input - innerCorner
			innerOffset := layerPos % (i - 1)

			res = numY + int(math.Abs(float64(innerOffset-numY)))
			break
		}
	}
	fmt.Printf("Part 1: %d\n", res)

	// part2 grid
	length := 12
	c := length / 2
	grid := make([][]int, length)
	for i := range grid {
		grid[i] = make([]int, length)
	}

	//grid[x][y]
	y, x := c, c
	grid[x][y] = 1
	sideLength := 2
	newLayer := false
	moved := 0
	direction := "left"

	// manually do first moves
	// go right 1
	x += 1
	//tLeft, mLeft, bLeft, tMid, bMid, tRight, mRight, bRight := grid[y+1][x-1], grid[y][x-1], grid[y-1][x-1], grid[y+1][x], grid[y-1][x], grid[y+1][x+1], grid[y][x+1], grid[y-1][x+1]
	tLeft, mLeft, bLeft, tMid, bMid, tRight, mRight, bRight := grid[x-1][y+1], grid[x-1][y], grid[x-1][y-1], grid[x][y+1], grid[x][y-1], grid[x+1][y+1], grid[x+1][y], grid[x+1][y-1]
	grid[x][y] = tLeft + mLeft + bLeft + tMid + bMid + tRight + mRight + bRight
	//fmt.Printf("Moved right setting %d at X: %d Y: %d \n", tLeft+mLeft+bLeft+tMid+bMid+tRight+mRight+bRight, x, y)

	// go up 1
	y -= 1
	tLeft, mLeft, bLeft, tMid, bMid, tRight, mRight, bRight = grid[x-1][y+1], grid[x-1][y], grid[x-1][y-1], grid[x][y+1], grid[x][y-1], grid[x+1][y+1], grid[x+1][y], grid[x+1][y-1]
	grid[x][y] = tLeft + mLeft + bLeft + tMid + bMid + tRight + mRight + bRight
	//fmt.Printf("Moved up setting %d at X: %d Y: %d \n", tLeft+mLeft+bLeft+tMid+bMid+tRight+mRight+bRight, x, y)

	for {
		if newLayer {
			direction = "up"
			sideLength++
			newLayer = false
			// manually do a right move
			x += 1
			tLeft, mLeft, bLeft, tMid, bMid, tRight, mRight, bRight := grid[x-1][y+1], grid[x-1][y], grid[x-1][y-1], grid[x][y+1], grid[x][y-1], grid[x+1][y+1], grid[x+1][y], grid[x+1][y-1]
			//fmt.Printf("Moved right setting %d at X: %d Y: %d \n", tLeft+mLeft+bLeft+tMid+bMid+tRight+mRight+bRight, x, y)
			grid[x][y] = tLeft + mLeft + bLeft + tMid + bMid + tRight + mRight + bRight
			continue
		}

		if direction == "up" {
			y -= 1
		} else if direction == "left" {
			x -= 1
		} else if direction == "down" {
			y += 1
		} else if direction == "right" {
			x += 1
		}

		if grid[x][y] != 0 {
			fmt.Println("MODIFYING ALREADY TOUCHED CELL")
			break
		}

		tLeft, mLeft, bLeft, tMid, bMid, tRight, mRight, bRight = grid[x-1][y+1], grid[x-1][y], grid[x-1][y-1], grid[x][y+1], grid[x][y-1], grid[x+1][y+1], grid[x+1][y], grid[x+1][y-1]
		grid[x][y] = tLeft + mLeft + bLeft + tMid + bMid + tRight + mRight + bRight

		//fmt.Printf("Moved %s setting %d at X: %d Y: %d \n", direction, tLeft+mLeft+bLeft+tMid+bMid+tRight+mRight+bRight, x, y)

		moved++

		if grid[x][y] > input {
			fmt.Printf("Part 2: %d\n", grid[x][y])
			break
		}

		// fmt.Printf("Moved %s %d of %d\n", direction, moved, sideLength)

		if moved == sideLength {
			moved = 0
			if direction == "up" {
				sideLength++
				direction = "left"
			} else if direction == "left" {
				direction = "down"
			} else if direction == "down" {
				direction = "right"
			} else if direction == "right" {
				newLayer = true
			}
		}
	}

	fmt.Println()

	for y := 0; y < length+1; y++ {
		if y == length {
			for x := 0; x < length; x++ {
				fmt.Printf("%d\t", x+1)
			}
			break
		}
		for x := 0; x < length; x++ {
			if x == length/2 && y == length/2 {
				fmt.Printf("C %d\t", grid[x][y])
				continue
			}
			fmt.Printf("%d\t", grid[x][y])
		}
		fmt.Println("Y")
	}
}
