package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	inputLines := getInput()
	_ = inputLines
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
