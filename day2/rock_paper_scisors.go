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

	dict := map[string]string{
		"A": "rock",
		"X": "rock",
		"B": "paper",
		"Y": "paper",
		"C": "scissors",
		"Z": "scissors",
	}

	winScores := map[string]int{
		"rock":     7,
		"paper":    8,
		"scissors": 9,
	}

	looseScores := map[string]int{
		"rock":     1,
		"paper":    2,
		"scissors": 3,
	}

	drawScores := map[string]int{
		"rock":     4,
		"paper":    5,
		"scissors": 6,
	}

	roundEndDict := map[string]string{
		"X": "loose",
		"Y": "draw",
		"Z": "win",
	}

	file, err := os.Open("input_data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	strategy := [][]string{}
	totalStrategyScore := 0
	totalStrategyScoreAlt := 0
	for scanner.Scan() {
		round := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		opMove, secondVal := round[0], round[1]

		var outcome string
		roundScore := 0
		switch {
		case dict[opMove] == dict[secondVal]:
			outcome = "draw"
			roundScore = drawScores[dict[secondVal]]
		case dict[opMove] == "rock" && dict[secondVal] == "paper":
			outcome = "win"
			roundScore = winScores[dict[secondVal]]
		case dict[opMove] == "rock" && dict[secondVal] == "scissors":
			outcome = "loose"
			roundScore = looseScores[dict[secondVal]]
		case dict[opMove] == "paper" && dict[secondVal] == "rock":
			outcome = "loose"
			roundScore = looseScores[dict[secondVal]]
		case dict[opMove] == "paper" && dict[secondVal] == "scissors":
			outcome = "win"
			roundScore = winScores[dict[secondVal]]
		case dict[opMove] == "scissors" && dict[secondVal] == "rock":
			outcome = "win"
			roundScore = winScores[dict[secondVal]]
		case dict[opMove] == "scissors" && dict[secondVal] == "paper":
			outcome = "loose"
			roundScore = looseScores[dict[secondVal]]
		}

		totalStrategyScore += roundScore

		// 2nd puzzle
		var myProjMove string
		roundScoreAlt := 0
		switch {
		case roundEndDict[secondVal] == "draw":
			myProjMove = dict[opMove]
			roundScoreAlt = drawScores[myProjMove]
		case roundEndDict[secondVal] == "win":
			switch dict[opMove] {
			case "rock":
				myProjMove = "paper"
			case "paper":
				myProjMove = "scissors"
			case "scissors":
				myProjMove = "rock"
			}
			roundScoreAlt = winScores[myProjMove]
		case roundEndDict[secondVal] == "loose":
			switch dict[opMove] {
			case "rock":
				myProjMove = "scissors"
			case "paper":
				myProjMove = "rock"
			case "scissors":
				myProjMove = "paper"
			}
			roundScoreAlt = looseScores[myProjMove]
		}

		totalStrategyScoreAlt += roundScoreAlt

		strategy = append(strategy, []string{opMove, secondVal, dict[opMove], dict[secondVal], outcome, strconv.Itoa(roundScore), roundEndDict[secondVal]})
	}

	fmt.Printf("%v", strategy)
	fmt.Println("\n\nTotal score for a strategy: ", totalStrategyScore)
	fmt.Println("Total score for an alternative strategy: ", totalStrategyScoreAlt)

}
