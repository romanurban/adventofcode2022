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
	inputLines := getInput()

	var forest [99][99]int
	var visibilityMap [99][99]bool

	for i, line := range inputLines {
		for j, char := range line {
			forest[i][j] = int(char - '0')
		}
	}

	// edges are always visible
	for i := 0; i < 99; i++ {
		visibilityMap[i][0] = true
		visibilityMap[i][98] = true
	}
	for j := 0; j < 99; j++ {
		visibilityMap[0][j] = true
		visibilityMap[98][j] = true
	}

	for i := 1; i < 99; i++ {
		for j := 1; j < 99; j++ {
			targetTree := forest[i][j]

			visibleN := true
			// go north
			for n := i - 1; n >= 0; n-- {
				if targetTree <= forest[n][j] { // hidden from the north
					visibleN = false
					break
				}
			}
			visibleW := true
			// go west
			for w := j - 1; w >= 0; w-- {
				if targetTree <= forest[i][w] { // hidden from the west
					visibleW = false
					break
				}
			}
			visibleS := true
			// go south
			for s := i + 1; s < 99; s++ {
				if targetTree <= forest[s][j] { // hidden from the south
					visibleS = false
					break
				}
			}
			visibleE := true
			// go east
			for e := j + 1; e < 99; e++ {
				if targetTree <= forest[i][e] { // hidden from the east
					visibleE = false
					break
				}
			}

			if visibleN || visibleW || visibleS || visibleE {
				visibilityMap[i][j] = true
			}
		}
	}

	y := string("\033[33m") // yellow
	r := string("\033[0m")  // reset
	for i, row := range forest {
		for j, ch := range row {
			if visibilityMap[i][j] {
				fmt.Print(y + strconv.Itoa(ch))
			} else {
				fmt.Print(r + strconv.Itoa(ch))
			}
		}
		fmt.Println(r)
	}

	totalVisible := 0
	for _, row := range visibilityMap {
		for _, ch := range row {
			if ch {
				totalVisible++
			}
		}
	}

	fmt.Println("Total visible trees:", totalVisible)

	// part2
	var scenicScoreMap [99][99]int
	//  edges will alway have worst scenic scores
	for i := 0; i < 99; i++ {
		scenicScoreMap[i][0] = 0
		scenicScoreMap[i][98] = 0
	}
	for j := 0; j < 99; j++ {
		scenicScoreMap[0][j] = 0
		scenicScoreMap[98][j] = 0
	}

	for i := 1; i < 99; i++ {
		for j := 1; j < 99; j++ {
			targetTree := forest[i][j]

			// go north
			scoreN := 1
			for n := i - 1; n >= 0; n-- {
				if targetTree <= forest[n][j] {
					break
				}
				if n != 0 { // edge not visible
					scoreN++
				}
			}
			// go west
			scoreW := 1
			for w := j - 1; w >= 0; w-- {
				if targetTree <= forest[i][w] {
					break
				}
				if w != 0 { // edge not visible
					scoreW++
				}
			}
			// go south
			scoreS := 1
			for s := i + 1; s < 99; s++ {
				if targetTree <= forest[s][j] {
					break
				}
				if s != 98 { // edge not visible
					scoreS++
				}
			}
			// go east
			scoreE := 1
			for e := j + 1; e < 99; e++ {
				if targetTree <= forest[i][e] {
					break
				}
				if e != 98 { // edge not visible
					scoreE++
				}
			}

			scenicScore := scoreN * scoreW * scoreS * scoreE
			scenicScoreMap[i][j] = scenicScore
		}
	}

	maxScenicScore := 0
	for _, row := range scenicScoreMap {
		for _, score := range row {
			if score > maxScenicScore {
				maxScenicScore = score
			}
		}
	}

	fmt.Println("Max scenic score:", maxScenicScore)

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
	}

	return inputLines
}
