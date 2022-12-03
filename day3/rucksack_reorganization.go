package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mpvl/unique"
)

func main() {
	file, err := os.Open("input_data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalIntersects := 0
	i := 1
	var groupsOfThree [][][]int
	var currGroupOfThree [][]int
	for scanner.Scan() {
		rucksackItems := strings.TrimSpace(scanner.Text())

		// part2 gathering
		translatedrucksackItems := translateToPriority(rucksackItems)
		if i%3 != 0 {
			currGroupOfThree = append(currGroupOfThree, translatedrucksackItems)
		} else {
			currGroupOfThree = append(currGroupOfThree, translatedrucksackItems)
			groupsOfThree = append(groupsOfThree, currGroupOfThree)
			currGroupOfThree = nil
		}
		i++

		rucksackItemCount := len(rucksackItems)
		compartmentOneItems := rucksackItems[:rucksackItemCount/2]
		compartmentTwoItems := rucksackItems[rucksackItemCount/2:]
		compartmentOneItemsTrn := translateToPriority(compartmentOneItems)
		compartmentTwoItemsTrn := translateToPriority(compartmentTwoItems)
		fmt.Println(compartmentOneItems, compartmentTwoItems)
		fmt.Println(compartmentOneItemsTrn, compartmentTwoItemsTrn)

		intersectItems := Intersection(compartmentOneItemsTrn, compartmentTwoItemsTrn)
		totalIntersects += sum(intersectItems)
		fmt.Println(intersectItems)
	}
	fmt.Println("Total sum of intersects: ", totalIntersects)

	// part2 continue
	totalIntersectsOfThree := 0
	for grpIdx := range groupsOfThree {
		intersectOfTwo := IntersectionNonUniq(groupsOfThree[grpIdx][0], groupsOfThree[grpIdx][1])
		intersectOfThree := IntersectionNonUniq(intersectOfTwo, groupsOfThree[grpIdx][2])
		totalIntersectsOfThree += intersectOfThree[0]
	}
	fmt.Println("Total sum of intersects in groups of three: ", totalIntersectsOfThree)
}

func translateToPriority(compartmentItems string) []int {
	const asciiShiftL = 96
	const asciiShiftU = 38
	const asciiIdx = 90

	var translated []int

	for _, ch := range compartmentItems {
		asciiCh := int(ch)
		if asciiCh > asciiIdx { // lowercase
			asciiCh = asciiCh - asciiShiftL
		} else { // uppercase
			asciiCh = asciiCh - asciiShiftU
		}
		translated = append(translated, asciiCh)
	}

	return translated
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
	unique.Ints(&c)
	return
}

func IntersectionNonUniq(a, b []int) (c []int) {
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

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
