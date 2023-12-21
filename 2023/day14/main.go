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

var roundRock = "O"
var squareRock = "#"
var space = "."

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
	} else {
		tempInput = example1
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
	movedRocks := moveRocks(parsed, 1)
	result := calculateLoad(movedRocks)

	return result
}

func calculateLoad(movedRocks [][]string) int {
	filter := func(s string) bool { return s == roundRock }

	result := 0
	for index, row := range movedRocks {
		rockCount := len(util.Filter(row, filter))
		distanceFromN := len(movedRocks) - index
		result += rockCount * distanceFromN
	}
	return result
}

func moveRocks(parsed [][]string, direction int) [][]string {
	movedRocks := util.CopySlice(parsed)

	for rowIndex, _ := range movedRocks {
		if (rowIndex == 0 && direction > 0) || (rowIndex == len(movedRocks) && direction < 0) {
			continue
		}

		var workingRowIndex int
		if direction > 0 {
			workingRowIndex = rowIndex
		}
		if direction < 0 {
			workingRowIndex = len(parsed) - rowIndex - 1
		}

		row := parsed[workingRowIndex]

		for columnIndex, column := range row {
			if column != roundRock {
				continue
			}

			movedRocks[workingRowIndex][columnIndex] = space

			block := false
			slideIndex := workingRowIndex
			for block != true {
				slideIndex += direction * -1
				block = (slideIndex == -1 && direction > 0) || (slideIndex == len(movedRocks) && direction < 0) || movedRocks[slideIndex][columnIndex] != space
			}

			movedRocks[slideIndex+direction][columnIndex] = roundRock
		}
	}
	return movedRocks
}

func part2(input string) int {

	parsed := parseInput(input)
	movedRocks := moveRocks(parsed, -1)

	for _, row := range parsed {
		fmt.Println(strings.Join(row, ""))
	}
	fmt.Println()
	for _, row := range movedRocks {
		fmt.Println(strings.Join(row, ""))
	}
	return 0
}

func parseInput(input string) (ans [][]string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, strings.Split(line, ""))
	}
	return ans
}
