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

	var items []int
	items = append(items,
		64, 89, 65, 95,
		76, 66, 74, 87, 70, 56, 51, 66,
		91, 60, 63,
		92, 61, 79, 97, 79,
		93, 54,
		60, 79, 92, 69, 88, 82, 70,
		64, 57, 73, 89, 55, 53,
		62)
	var monkeyItems [8][]int
	var monkeyInspections [8]int
	fmt.Println(" ")
	fmt.Println("Initial levels: ")
	fmt.Println(items[:4])
	fmt.Println(items[4:12])
	fmt.Println(items[12:15])
	fmt.Println(items[15:20])
	fmt.Println(items[20:22])
	fmt.Println(items[22:29])
	fmt.Println(items[29:35])
	fmt.Println(items[35])
	monkeyItems[0] = []int{0, 1, 2, 3}
	monkeyItems[1] = []int{4, 5, 6, 7, 8, 9, 10, 11}
	monkeyItems[2] = []int{12, 13, 14}
	monkeyItems[3] = []int{15, 16, 17, 18, 19}
	monkeyItems[4] = []int{20, 21}
	monkeyItems[5] = []int{22, 23, 24, 25, 26, 27, 28}
	monkeyItems[6] = []int{29, 30, 31, 32, 33, 34}
	monkeyItems[7] = []int{35}

	fmt.Println("Monkey items:", monkeyItems)

	for round := range [20]int{} {
		_ = round
		fmt.Println("Round:", round)
		// monkey 0 turn
		for i, itemIndex := range monkeyItems[0] {
			_ = i
			monkeyInspections[0]++
			// inspect
			items[itemIndex] = items[itemIndex] * 7
			// bored
			items[itemIndex] = int(math.Floor(float64(items[itemIndex] / 3)))
			// test and pass
			if items[itemIndex]%3 == 0 {
				monkeyItems[4] = append(monkeyItems[4], itemIndex)
			} else {
				monkeyItems[1] = append(monkeyItems[1], itemIndex)
			}
			monkeyItems[0] = remove(monkeyItems[0], itemIndex)
		}
		// monkey 1 turn
		for i, itemIndex := range monkeyItems[1] {
			_ = i
			monkeyInspections[1]++
			// inspect
			items[itemIndex] = items[itemIndex] + 5
			// bored
			items[itemIndex] = int(math.Floor(float64(items[itemIndex] / 3)))
			// test and pass
			if items[itemIndex]%13 == 0 {
				monkeyItems[7] = append(monkeyItems[7], itemIndex)
			} else {
				monkeyItems[3] = append(monkeyItems[3], itemIndex)
			}
			monkeyItems[1] = remove(monkeyItems[1], itemIndex)
		}
		// monkey 2 turn
		for i, itemIndex := range monkeyItems[2] {
			_ = i
			monkeyInspections[2]++
			// inspect
			items[itemIndex] = items[itemIndex] * items[itemIndex]
			// bored
			items[itemIndex] = int(math.Floor(float64(items[itemIndex] / 3)))
			// test and pass
			if items[itemIndex]%2 == 0 {
				monkeyItems[6] = append(monkeyItems[6], itemIndex)
			} else {
				monkeyItems[5] = append(monkeyItems[5], itemIndex)
			}
			monkeyItems[2] = remove(monkeyItems[2], itemIndex)
		}
		// monkey 3 turn
		for i, itemIndex := range monkeyItems[3] {
			_ = i
			monkeyInspections[3]++
			// inspect
			items[itemIndex] = items[itemIndex] + 6
			// bored
			items[itemIndex] = int(math.Floor(float64(items[itemIndex] / 3)))
			// test and pass
			if items[itemIndex]%11 == 0 {
				monkeyItems[2] = append(monkeyItems[2], itemIndex)
			} else {
				monkeyItems[6] = append(monkeyItems[6], itemIndex)
			}
			monkeyItems[3] = remove(monkeyItems[3], itemIndex)
		}
		// monkey 4 turn
		for i, itemIndex := range monkeyItems[4] {
			_ = i
			monkeyInspections[4]++
			// inspect
			items[itemIndex] = items[itemIndex] * 11
			// bored
			items[itemIndex] = int(math.Floor(float64(items[itemIndex] / 3)))
			// test and pass
			if items[itemIndex]%5 == 0 {
				monkeyItems[1] = append(monkeyItems[1], itemIndex)
			} else {
				monkeyItems[7] = append(monkeyItems[7], itemIndex)
			}
			monkeyItems[4] = remove(monkeyItems[4], itemIndex)
		}
		// monkey 5 turn
		for i, itemIndex := range monkeyItems[5] {
			_ = i
			monkeyInspections[5]++
			// inspect
			items[itemIndex] = items[itemIndex] + 8
			// bored
			items[itemIndex] = int(math.Floor(float64(items[itemIndex] / 3)))
			// test and pass
			if items[itemIndex]%17 == 0 {
				monkeyItems[4] = append(monkeyItems[4], itemIndex)
			} else {
				monkeyItems[0] = append(monkeyItems[0], itemIndex)
			}
			monkeyItems[5] = remove(monkeyItems[5], itemIndex)
		}
		// monkey 6 turn
		for i, itemIndex := range monkeyItems[6] {
			_ = i
			monkeyInspections[6]++
			// inspect
			items[itemIndex] = items[itemIndex] + 1
			// bored
			items[itemIndex] = int(math.Floor(float64(items[itemIndex] / 3)))
			// test and pass
			if items[itemIndex]%19 == 0 {
				monkeyItems[0] = append(monkeyItems[0], itemIndex)
			} else {
				monkeyItems[5] = append(monkeyItems[5], itemIndex)
			}
			monkeyItems[6] = remove(monkeyItems[6], itemIndex)
		}
		// monkey 7 turn
		for i, itemIndex := range monkeyItems[7] {
			_ = i
			monkeyInspections[7]++
			// inspect
			items[itemIndex] = items[itemIndex] + 4
			// bored
			items[itemIndex] = int(math.Floor(float64(items[itemIndex] / 3)))
			// test and pass
			if items[itemIndex]%7 == 0 {
				monkeyItems[3] = append(monkeyItems[3], itemIndex)
			} else {
				monkeyItems[2] = append(monkeyItems[2], itemIndex)
			}
			monkeyItems[7] = remove(monkeyItems[7], itemIndex)
		}
		fmt.Println("Monkey items after " + strconv.Itoa(round) + " round")
		fmt.Println(monkeyItems)
		for monkey := range monkeyItems {
			fmt.Println("Monkey ", monkey, "holds items")
			for _, idx := range monkeyItems[monkey] {
				fmt.Print(items[idx], " ")
			}
			fmt.Println(" ")
		}
		fmt.Println("Levels after " + strconv.Itoa(round) + " round")
		fmt.Println(items)
		fmt.Println("Monkey total inspections: ", monkeyInspections)
	}
}

func remove(s []int, val int) []int {
	i := indexOf(s, val)
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func indexOf(s []int, val int) int {
	for k, v := range s {
		if val == v {
			return k
		}
	}
	return -1 //not found.
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
