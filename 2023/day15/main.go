package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"advent-of-code-go/util"
)

//go:embed input.txt
var input string

//go:embed example1.txt
var example1 string

//go:embed example2.txt
var example2 string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}

	example1 = strings.TrimRight(example1, "\n")
	if len(input) == 0 {
		panic("empty example1.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	var inputType string
	flag.StringVar(&inputType, "inputType", "puzzle", "puzzle or example1")
	flag.Parse()

	var tempInput string
	if inputType == "puzzle" {
		tempInput = input
	} else if inputType == "example1" {
		tempInput = example1
	} else if inputType == "example2" {
		tempInput = example2
	}

	if part == 1 {
		ans := part1(tempInput)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(tempInput)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {

	parsed := parseInput(input)
	result := 0
	for _, value := range parsed {
		result += getInitialisationStep(value)
	}

	return result
}

func getInitialisationStep(input string) int {
	result := 0
	for i := 0; i < len(input); i++ {
		ascii := int(input[i])
		result += ascii
		result = result * 17
		result = result % 256
	}
	return result
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (ans []string) {
	return strings.Split(input, ",")
}
