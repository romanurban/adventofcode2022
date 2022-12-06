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
	file, err := os.Open("input_data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	type void struct{}
	var member void

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fmt.Println(line)

		k := 0
		n := 4
		l := 0 // starting marker for 2nd puzzle
		for i := 0; n < len(line); i++ {
			set := make(map[rune]void)
			for j := k; j < n; j++ {
				set[rune(line[j])] = member
			}
			if len(set) == 4 {
				for i := range set {
					fmt.Println(strconv.QuoteRune(i))
				}
				fmt.Println("Packet marker at", i+4)
				n = len(line)
				l = i
				break
			}
			k++
			n++
		}
		// 2nd puzzle
		k = l
		n = l + 14
		for i := l; n < len(line); i++ {
			set := make(map[rune]void)
			for j := k; j < n; j++ {
				set[rune(line[j])] = member
			}
			if len(set) == 14 {
				for i := range set {
					fmt.Println(strconv.QuoteRune(i))
				}
				fmt.Println("Start-of-message marker at", i+14)
				n = len(line)
				l = i
				break
			}
			k++
			n++
		}
	}
}
