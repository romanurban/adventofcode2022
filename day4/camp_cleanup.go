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
	// [1-99] array of all camp sections
	var campSections = [99]int{}
	for i := 0; i < len(campSections); i++ {
		campSections[i] = i + 1
	}

	file, err := os.Open("input_data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pairsWithRangeIntersect := 0
	pairsIntersect := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		groupRanges := strings.Split(line, ",")
		groupFrstRangeBoundaries := strings.Split(groupRanges[0], "-")
		groupScndRangeBoundaries := strings.Split(groupRanges[1], "-")
		groupFrstRangeBoundaryL, _ := strconv.Atoi(groupFrstRangeBoundaries[0])
		groupFrstRangeBoundaryH, _ := strconv.Atoi(groupFrstRangeBoundaries[1])
		groupScndRangeBoundaryL, _ := strconv.Atoi(groupScndRangeBoundaries[0])
		groupScndRangeBoundaryH, _ := strconv.Atoi(groupScndRangeBoundaries[1])
		groupFrstSlice := campSections[groupFrstRangeBoundaryL-1 : groupFrstRangeBoundaryH]
		groupScndSlice := campSections[groupScndRangeBoundaryL-1 : groupScndRangeBoundaryH]

		intersect := Intersection(groupFrstSlice, groupScndSlice)

		if len(intersect) == len(groupFrstSlice) || len(intersect) == len(groupScndSlice) {
			pairsWithRangeIntersect++
		}

		if len(intersect) > 0 {
			pairsIntersect++
		}

		fmt.Println(line)
		fmt.Println(groupFrstSlice, groupScndSlice)
	}

	fmt.Println("Num of pairs with range full overlap", pairsWithRangeIntersect)
	fmt.Println("Num of pairs intersect", pairsIntersect)
}

func Intersection(a, b []int) (c []int) {
	m := make(map[int]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}
