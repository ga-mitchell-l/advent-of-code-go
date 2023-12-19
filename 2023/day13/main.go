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

	for i, pattern := range parsed {
		_ = i
		for j, row := range pattern {
			_ = j

			for k, char := range row {
				_ = char

				if k == 0 {
					continue
				}

				nextChar := row[k-1]
				_ = nextChar

			}

		}
	}

	// To find the reflection in each pattern, you need to find a perfect reflection
	// across either a horizontal line between two rows or across a vertical line between two columns.

	return 0
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (ans [][][]string) {

	patterns := strings.Split(input, "\n\n")

	processedPatterns := make([][][]string, len(patterns))
	for i, pattern := range patterns {
		rows := strings.Split(pattern, "\n")
		processedPatterns[i] = make([][]string, len(rows))

		for j, row := range rows {
			processedPatterns[i][j] = strings.Split(row, "")
		}

	}

	return processedPatterns
}
