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
	_ = inputLines

	var commandBuffer []string
	registerX := 1
	_ = registerX

	for _, line := range inputLines {
		if line == "noop" {
			commandBuffer = append(commandBuffer, line)
		} else {
			addx := strings.Split(line, " ")
			addx_f := addx[0] + "_f:" + addx[1]
			addx_s := addx[0] + "_s:" + addx[1]
			commandBuffer = append(commandBuffer, addx_f, addx_s)
		}
	}
	fmt.Println("Transformed command buffer:")
	for i, command := range commandBuffer {
		fmt.Println(i+1, command)
	}

	totalStrength := 0
	for i := 20; i <= 220; i += 40 {
		signalStrength := getSignalStrength(commandBuffer, i-1)
		fmt.Println("signalStrength on level", i, "=", signalStrength, "result:", signalStrength*i)
		signalStrength = signalStrength * i
		totalStrength += signalStrength
	}
	fmt.Println("Total strength", totalStrength)

	// part 2
	var row1 string
	var row2 string
	var row3 string
	var row4 string
	var row5 string
	var row6 string

	for i := 0; i < 40; i++ {
		pos := i
		signalStrength := getSignalStrength(commandBuffer, i)
		fmt.Println(pos, "[", signalStrength-1, signalStrength, signalStrength+1, "]")
		if pos >= signalStrength-1 && pos <= signalStrength+1 {
			row1 += "#"
		} else {
			row1 += "."
		}
	}
	for i := 40; i < 80; i++ {
		pos := i - 40
		signalStrength := getSignalStrength(commandBuffer, i)
		fmt.Println(pos, "[", signalStrength-1, signalStrength, signalStrength+1, "]")
		if pos >= signalStrength-1 && pos <= signalStrength+1 {
			row2 += "#"
		} else {
			row2 += "."
		}
	}
	for i := 80; i < 120; i++ {
		pos := i - 80
		signalStrength := getSignalStrength(commandBuffer, i)
		fmt.Println(pos, "[", signalStrength-1, signalStrength, signalStrength+1, "]")
		if pos >= signalStrength-1 && pos <= signalStrength+1 {
			row3 += "#"
		} else {
			row3 += "."
		}
	}
	for i := 120; i < 160; i++ {
		pos := i - 120
		signalStrength := getSignalStrength(commandBuffer, i)
		fmt.Println(pos, "[", signalStrength-1, signalStrength, signalStrength+1, "]")
		if pos >= signalStrength-1 && pos <= signalStrength+1 {
			row4 += "#"
		} else {
			row4 += "."
		}
	}
	for i := 160; i < 200; i++ {
		pos := i - 160
		signalStrength := getSignalStrength(commandBuffer, i)
		fmt.Println(pos, "[", signalStrength-1, signalStrength, signalStrength+1, "]")
		if pos >= signalStrength-1 && pos <= signalStrength+1 {
			row5 += "#"
		} else {
			row5 += "."
		}
	}
	for i := 200; i < 240; i++ {
		pos := i - 200
		signalStrength := getSignalStrength(commandBuffer, i)
		fmt.Println(pos, "[", signalStrength-1, signalStrength, signalStrength+1, "]")
		if pos >= signalStrength-1 && pos <= signalStrength+1 {
			row6 += "#"
		} else {
			row6 += "."
		}
	}

	fmt.Println(len(row1), row1)
	fmt.Println(len(row2), row2)
	fmt.Println(len(row3), row3)
	fmt.Println(len(row4), row4)
	fmt.Println(len(row5), row5)
	fmt.Println(len(row6), row6)
}

func getSignalStrength(commandBuffer []string, cycle int) int {
	addxUntil := 0
	strength := 1
	if commandBuffer[cycle] == "noop" {
		addxUntil = cycle
	} else if strings.HasPrefix(commandBuffer[cycle], "addx_s") {
		addxUntil = cycle - 1
	} else {
		addxUntil = cycle
	}

	for i := 0; i < addxUntil; i++ {
		if commandBuffer[i] == "noop" {
			continue
		} else if strings.HasPrefix(commandBuffer[i], "addx_s") {
			command := strings.Split(commandBuffer[i], ":")
			V, _ := strconv.Atoi(command[1])
			strength += V
		}
	}
	return strength
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
