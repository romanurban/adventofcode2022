package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// DOESN'T WORK
func main() {
	inputLines := getInput()
	_ = inputLines

	var grid [999][999]rune
	var headVisitGrid [999][999]rune
	var knot1VisitGrid [999][999]rune
	var knot2VisitGrid [999][999]rune
	var knot3VisitGrid [999][999]rune
	var knot4VisitGrid [999][999]rune
	var knot5VisitGrid [999][999]rune
	var knot6VisitGrid [999][999]rune
	var knot7VisitGrid [999][999]rune
	var knot8VisitGrid [999][999]rune
	var knot9VisitGrid [999][999]rune
	// put starting point in the center
	var xH, x1, x2, x3, x4, x5, x6, x7, x8, x9 int
	var yH, y1, y2, y3, y4, y5, y6, y7, y8, y9 int
	xH, x1, x2, x3, x4, x5, x6, x7, x8, x9 = 499, 499, 499, 499, 499, 499, 499, 499, 499, 499
	yH, y1, y2, y3, y4, y5, y6, y7, y8, y9 = 499, 499, 499, 499, 499, 499, 499, 499, 499, 499
	headVisitGrid[yH][xH] = 's'
	knot1VisitGrid[y1][x1] = 's'
	knot2VisitGrid[y2][x2] = 's'
	knot3VisitGrid[y3][x3] = 's'
	knot4VisitGrid[y4][x4] = 's'
	knot5VisitGrid[y5][x5] = 's'
	knot6VisitGrid[y6][x6] = 's'
	knot7VisitGrid[y7][x7] = 's'
	knot8VisitGrid[y8][x8] = 's'
	knot9VisitGrid[y9][x9] = 's'

	for _, line := range inputLines {
		move := strings.Split(line, " ")
		direction := move[0]
		steps, _ := strconv.Atoi(move[1])

		for i := 1; i <= steps; i++ {
			switch direction {
			case "U":
				yH--
			case "D":
				yH++
			case "R":
				xH++
			case "L":
				xH--
			}

			if direction == "L" && steps == 17 {
				_ = i
			}
			headVisitGrid[yH][xH] = 'H'
			y1, x1 = knotMove(yH, xH, y1, x1, direction)
			knot1VisitGrid[y1][x1] = '1'
			y2, x2 = knotMove(y1, x1, y2, x2, direction)
			knot2VisitGrid[y2][x2] = '2'
			y3, x3 = knotMove(y2, x2, y3, x3, direction)
			knot3VisitGrid[y3][x3] = '3'
			y4, x4 = knotMove(y3, x3, y4, x4, direction)
			knot4VisitGrid[y4][x4] = '4'
			y5, x5 = knotMove(y4, x4, y5, x5, direction)
			knot5VisitGrid[y5][x5] = '5'
			y6, x6 = knotMove(y5, x5, y6, x6, direction)
			knot6VisitGrid[y6][x6] = '6'
			y7, x7 = knotMove(y6, x6, y7, x7, direction)
			knot7VisitGrid[y7][x7] = '7'
			y8, x8 = knotMove(y7, x7, y8, x8, direction)
			knot8VisitGrid[y8][x8] = '8'
			y9, x9 = knotMove(y8, x8, y9, x9, direction)
			knot9VisitGrid[y9][x9] = '9'

			for i := 0; i < 999; i++ {
				for j := 0; j < 999; j++ {
					grid[i][j] = rune(0)
				}
			}
			grid[y9][x9] = '9'
			grid[y8][x8] = '8'
			grid[y7][x7] = '7'
			grid[y6][x6] = '6'
			grid[y5][x5] = '5'
			grid[y4][x4] = '4'
			grid[y3][x3] = '3'
			grid[y2][x2] = '2'
			grid[y1][x1] = '1'
			grid[yH][xH] = 'H'
		}
	}
	printTotalVisited(headVisitGrid, 0)
	printTotalVisited(knot1VisitGrid, 1)
	printTotalVisited(knot2VisitGrid, 2)
	printTotalVisited(knot3VisitGrid, 3)
	printTotalVisited(knot4VisitGrid, 4)
	printTotalVisited(knot5VisitGrid, 5)
	printTotalVisited(knot6VisitGrid, 6)
	printTotalVisited(knot7VisitGrid, 7)
	printTotalVisited(knot8VisitGrid, 8)
	printTotalVisited(knot9VisitGrid, 9)
}

// some kind of a bug with movement logic
func knotMove(y int, x int, yT int, xT int, direction string) (int, int) {
	if x == xT || y == yT { // if in the same row or column
		if math.Abs(float64(x-xT)) == 1 || math.Abs(float64(y-yT)) == 1 {
			// adjacent cells, tail doesn't move
			return yT, xT
		}
		if x == xT && y == yT {
			// H and T overlap
			return yT, xT
		}

		switch direction {
		case "U":
			yT--
		case "D":
			yT++
		case "R":
			xT++
		case "L":
			xT--
		}
	} else { // if diagonal position
		if math.Abs(float64(x-xT)) == 1 && math.Abs(float64(y-yT)) == 1 {
			// adjacent cells, tail doesn't move
			return yT, xT
		}
		// top left
		if y < yT && x < xT {
			yT--
			xT--
		}

		// top right
		if y < yT && x > xT {
			yT--
			xT++
		}

		// bottom left
		if y > yT && x < xT {
			yT++
			xT--
		}

		// bottom right
		if y > yT && x > xT {
			yT++
			xT++
		}
	}
	return yT, xT
}

func printTotalVisited(grid [999][999]rune, knot int) {
	totalVisited := 0
	for i, row := range grid {
		for j, _ := range row {
			if grid[i][j] != rune(0) {
				totalVisited++
			}
		}
	}
	fmt.Println("Positions rope's " + strconv.Itoa(knot) + "th knot visited: " + strconv.Itoa(totalVisited))
}

func getInput() []string {
	file, err := os.Open("input_data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var inputLines []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		inputLines = append(inputLines, line)
		//fmt.Println(line)
	}

	return inputLines
}
