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
	inputLines = inputLines[1:]

	var path string
	dirTreeSizes := make(map[string]int)
	prevPath := ""
	dirDirectSize := 0
	_ = prevPath
	for _, line := range inputLines {
		// go up
		if line == "$ cd .." {
			path = path[:strings.LastIndex(path, "/")]
			if len(path) == 0 {
				path = "/"
			}
			fmt.Println(path)
			continue
		}

		// go down
		if strings.HasPrefix(line, "$ cd") {
			if path == "/" {
				path = ""
			}
			path = path + "/" + line[5:]
			prevPath = path
			fmt.Println(path)
			continue
		}

		if line[0] >= '0' && line[0] <= '9' { // file met
			str := strings.Split(line, " ")[0]
			fileSize, _ := strconv.Atoi(str)
			dirDirectSize += fileSize
			dirTreeSizes[path] += fileSize
		} else { // catalog met
			fileSize := 0
			dirTreeSizes[path] += fileSize
		}

	}

	fmt.Println("##############################################")
	fmt.Println(dirTreeSizes)

	indTreeSizes := make(map[string]int)

	for key := range dirTreeSizes {
		for path := range dirTreeSizes {
			if strings.HasPrefix(path, key) {
				indTreeSizes[key] += dirTreeSizes[path]
			}
		}
	}

	fmt.Println("##############################################")
	fmt.Println(indTreeSizes)

	part1Total := 0
	for key := range indTreeSizes {
		if indTreeSizes[key] < 100000 {
			part1Total += indTreeSizes[key]
		}
	}

	fmt.Println("##############################################")
	fmt.Println("The sum of the total for dirs <100K: ", part1Total)

	// part 2
	totalSpace := 70000000
	reqForUpdate := 30000000
	totalUsed := indTreeSizes[""]
	totalFree := totalSpace - totalUsed
	reqToCleanup := reqForUpdate - (totalSpace - totalUsed)

	fmt.Println("############## PART TWO ################")
	fmt.Println("Total used of 70000000: ", totalUsed)
	fmt.Println("Total free (min 30000000): ", totalFree)
	fmt.Println("For a cleanup at least: ", reqToCleanup)

	cleanupCandidateSpace := totalSpace
	cleanupCandidatePath := ""
	for key := range indTreeSizes {
		if indTreeSizes[key] > reqToCleanup {
			if indTreeSizes[key] < cleanupCandidateSpace {
				cleanupCandidateSpace = indTreeSizes[key]
				cleanupCandidatePath = key
			}
		}
	}
	fmt.Println("Best to delete: ", cleanupCandidatePath)
	fmt.Println("Of total size: ", cleanupCandidateSpace)

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
