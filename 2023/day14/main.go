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
	_ = parsed

	movedRocks := moveRocks(parsed)
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

func moveRocks(parsed [][]string) [][]string {
	movedRocks := util.CopySlice(parsed)
	for rowIndex, row := range movedRocks {
		if rowIndex == 0 {
			continue
		}

		for columnIndex, column := range row {
			if column != roundRock {
				continue
			}

			movedRocks[rowIndex][columnIndex] = space

			block := false
			slideIndex := rowIndex
			for block != true {
				slideIndex--
				block = slideIndex == -1 || movedRocks[slideIndex][columnIndex] != space
			}

			movedRocks[slideIndex+1][columnIndex] = roundRock
		}
	}
	return movedRocks
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (ans [][]string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, strings.Split(line, ""))
	}
	return ans
}
