package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input_data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	elf_items := make(map[string][]int)

	key_part := "elf_"
	i := 0
	var current_key = key_part + strconv.Itoa(i)
	for scanner.Scan() {
		if len(strings.TrimSpace(scanner.Text())) != 0 {
			curr_item, _ := strconv.Atoi(scanner.Text())
			elf_items[current_key] = append(elf_items[current_key], curr_item)
		} else {
			i++
			current_key = key_part + strconv.Itoa(i)
		}
	}

	max_calories := 0
	var elf_items_calories []int
	var elf_max_calories string
	for elf, cur_elf_items := range elf_items {
		curr_elf_calories := 0
		fmt.Println("Elf:", elf, "=>", "Calories of items:", cur_elf_items)

		for i := 0; i < len(cur_elf_items); i++ {
			curr_elf_calories = curr_elf_calories + cur_elf_items[i]
		}

		elf_items_calories = append(elf_items_calories, curr_elf_calories)

		if curr_elf_calories > max_calories {
			max_calories = curr_elf_calories
			elf_max_calories = elf
		}

		fmt.Println("Elf:", elf, "=>", "Total calories:", curr_elf_calories)
	}

	fmt.Println("Elf with max calories:", elf_max_calories, "=>", "Carries total calories:", max_calories)

	sort.Sort(sort.Reverse(sort.IntSlice(elf_items_calories)))

	sum_top_3_elfs := elf_items_calories[0] + elf_items_calories[1] + elf_items_calories[2]

	fmt.Println("Top 3 Elfs carry total calories :", sum_top_3_elfs)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
