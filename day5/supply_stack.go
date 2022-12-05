package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
)

func main() {
	file, err := os.Open("input_data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	supplyStacks := initSupplyStacks()
	supplyStacks90001 := initSupplyStacks()

	printTopStacksCrates(supplyStacks)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "move") {
			cargoCraneMoveLine := strings.Replace(line, "move ", "", -1)
			cargoCraneMoveLine = strings.Replace(cargoCraneMoveLine, " from ", " ", -1)
			cargoCraneMoveLine = strings.Replace(cargoCraneMoveLine, " to ", " ", -1)
			cargoCraneMove := strings.Split(cargoCraneMoveLine, " ")
			// fmt.Println(cargoCraneMove)
			craneCount, _ := strconv.Atoi(cargoCraneMove[0])
			performMove(supplyStacks, craneCount, cargoCraneMove[1], cargoCraneMove[2])
			performMove90001(supplyStacks90001, craneCount, cargoCraneMove[1], cargoCraneMove[2])
		}
	}

	printTopStacksCrates(supplyStacks)
	printTopStacksCrates(supplyStacks90001)
}

func printTopStacksCrates(supplyStacks map[string]*stack.Stack) {
	fmt.Println("Stack top crates: ")
	fmt.Print("[")
	for i := 1; i < 10; i++ {
		iStr := strconv.Itoa(i)
		fmt.Print(supplyStacks["stack"+iStr].Peek())
	}
	fmt.Println("]")
}

func initSupplyStacks() map[string]*stack.Stack {
	supplyStacks := make(map[string]*stack.Stack)
	stack1 := stack.New()
	pushCratesToStack(stack1, []string{"R", "G", "H", "Q", "S", "B", "T", "N"})
	supplyStacks["stack1"] = stack1
	stack2 := stack.New()
	pushCratesToStack(stack2, []string{"H", "S", "F", "D", "P", "Z", "J"})
	supplyStacks["stack2"] = stack2
	stack3 := stack.New()
	pushCratesToStack(stack3, []string{"Z", "H", "V"})
	supplyStacks["stack3"] = stack3
	stack4 := stack.New()
	pushCratesToStack(stack4, []string{"M", "Z", "J", "F", "G", "H"})
	supplyStacks["stack4"] = stack4
	stack5 := stack.New()
	pushCratesToStack(stack5, []string{"T", "Z", "C", "D", "L", "M", "S", "R"})
	supplyStacks["stack5"] = stack5
	stack6 := stack.New()
	pushCratesToStack(stack6, []string{"M", "T", "W", "V", "H", "Z", "J"})
	supplyStacks["stack6"] = stack6
	stack7 := stack.New()
	pushCratesToStack(stack7, []string{"T", "F", "P", "L", "Z"})
	supplyStacks["stack7"] = stack7
	stack8 := stack.New()
	pushCratesToStack(stack8, []string{"Q", "V", "W", "S"})
	supplyStacks["stack8"] = stack8
	stack9 := stack.New()
	pushCratesToStack(stack9, []string{"W", "H", "L", "M", "T", "D", "N", "C"})
	supplyStacks["stack9"] = stack9

	return supplyStacks
}

func pushCratesToStack(stack *stack.Stack, elems []string) {
	for _, elem := range elems {
		stack.Push(elem)
	}
}

func popCratesFromStack(stack *stack.Stack, count int) []string {
	var elems []string
	for i := 0; i < count; i++ {
		elems = append(elems, stack.Pop().(string))
	}
	return elems
}

func performMove(supplyStacks map[string]*stack.Stack, count int, from string, to string) {
	poppedStacks := popCratesFromStack(supplyStacks["stack"+from], count)
	pushCratesToStack(supplyStacks["stack"+to], poppedStacks)
}

func performMove90001(supplyStacks map[string]*stack.Stack, count int, from string, to string) {
	poppedStacks := popCratesFromStack(supplyStacks["stack"+from], count)
	reverse(poppedStacks)
	pushCratesToStack(supplyStacks["stack"+to], poppedStacks)
}

func reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}
