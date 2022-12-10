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

func main() {
	inputLines := getInput()
	_ = inputLines

	var tailVisitGrid [999][999]rune
	var grid [999][999]rune
	// put starting point in the center
	x := 499
	y := 499
	xT := x
	yT := y
	tailVisitGrid[yT][xT] = 's'
	grid[y][x] = 's'

	// perform a first move [U 1]
	y = y - 1
	grid[y][x] = 'H'

	for i, line := range inputLines {
		// skip the first move as it already performed
		if i == 0 {
			continue
		}

		move := strings.Split(line, " ")
		fmt.Println(move)
		direction := move[0]
		steps, _ := strconv.Atoi(move[1])

		for i := 1; i <= steps; i++ {
			switch direction {
			case "U":
				y--
			case "D":
				y++
			case "R":
				x++
			case "L":
				x--
			}

			grid[y][x] = 'H'
			moveTail(&tailVisitGrid, &y, &x, &yT, &xT, direction)
		}
	}

	/* // visualisation
	g := string("\033[32m") // green
	w := string("\033[33m") // yellow
	r := string("\033[0m")  // reset
	for i, row := range grid {
		for j, col := range row {
			if i == y && j == x {
				fmt.Print(w + string(col))
			} else if grid[i][j] != rune(0) {
				fmt.Print(g + string(col))
			} else {
				fmt.Print(r + ".")
			}
		}
		fmt.Println(r)
	}
	fmt.Println(r)
	for i, row := range tailVisitGrid {
		for j, col := range row {
			if tailVisitGrid[i][j] != rune(0) {
				fmt.Print(g + string(col))
			} else {
				fmt.Print(r + ".")
			}
		}
		fmt.Println(r)
	}*/
	totalHeadVisited := 0
	for i, row := range grid {
		for j, _ := range row {
			if grid[i][j] != rune(0) {
				totalHeadVisited++
			}
		}
	}
	fmt.Println("Positions rope head visited:", totalHeadVisited)

	totalTailVisited := 0
	for i, row := range tailVisitGrid {
		for j, _ := range row {
			if tailVisitGrid[i][j] != rune(0) {
				totalTailVisited++
			}
		}
	}
	fmt.Println("Positions rope tail visited:", totalTailVisited)

}

func moveTail(tailVisitGrid *[999][999]rune, y *int, x *int, yT *int, xT *int, direction string) {
	if *x == *xT || *y == *yT { // if in the same row or column
		if *x == *xT && *y == *yT {
			// H and T overlap
			return
		}

		if math.Abs(float64(*x-*xT)) == 1 || math.Abs(float64(*y-*yT)) == 1 {
			// adjacent cells, tail doesn't move
			return
		}
		switch direction {
		case "U":
			*yT--
		case "D":
			*yT++
		case "R":
			*xT++
		case "L":
			*xT--
		}
		tailVisitGrid[*yT][*xT] = 'T'
	} else { // if diagonal position
		if math.Abs(float64(*x-*xT)) == 1 && math.Abs(float64(*y-*yT)) == 1 {
			// adjacent cells, tail doesn't move
			return
		}
		// top left
		if *y < *yT && *x < *xT {
			*yT--
			*xT--
		}

		// top right
		if *y < *yT && *x > *xT {
			*yT--
			*xT++
		}

		// bottom left
		if *y > *yT && *x < *xT {
			*yT++
			*xT--
		}

		// bottom right
		if *y > *yT && *x > *xT {
			*yT++
			*xT++
		}

		tailVisitGrid[*yT][*xT] = 'T'
	}

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
		fmt.Println(line)
	}

	return inputLines
}
