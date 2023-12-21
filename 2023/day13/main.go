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

	result := 0
	for _, pattern := range parsed {
		transposedPattern := util.Transpose(pattern)

		verticalMirrors := getMirrorsByPattern(pattern, 0)
		horizontalMirrors := getMirrorsByPattern(transposedPattern, 0)

		result = summarizePattern(verticalMirrors, horizontalMirrors, result)
	}

	return result
}

func summarizePattern(verticalMirrors []int, horizontalMirrors []int, result int) int {
	// add up the number of columns to the left of each vertical line of reflection
	pvsum := 0
	for _, val := range verticalMirrors {
		pvsum += val + 1
	}

	//  also add 100 multiplied by the number of rows above each horizontal line of reflection
	phsum := 0
	for _, val := range horizontalMirrors {
		phsum += 100 * (val + 1)
	}
	result += pvsum + phsum
	return result
}

func isColumnMirror(k int, firstRow []string) bool {
	mirrorLeftIndex := k
	mirrorRightIndex := k + 1
	mirror := firstRow[mirrorLeftIndex] == firstRow[mirrorRightIndex]

	for mirrorLeftIndex > 0 && mirrorRightIndex < len(firstRow)-1 && mirror {
		mirrorLeftIndex--
		mirrorRightIndex++
		mirror = firstRow[mirrorLeftIndex] == firstRow[mirrorRightIndex]
	}
	return mirror
}

func part2(input string) int {
	parsed := parseInput(input)
	result := 0

	for _, pattern := range parsed {
		transposedPattern := util.Transpose(pattern)
		verticalMirrors := getMirrorsByPattern(pattern, 1)
		horizontalMirrors := getMirrorsByPattern(transposedPattern, 1)

		result = summarizePattern(verticalMirrors, horizontalMirrors, result)
	}

	return result
}

func getMirrorsByPattern(pattern [][]string, matchDiff int) []int {
	var patternMirrors []int
	patternMirrorSet := make(map[int]int)
	rowCount := len(pattern)

	for _, row := range pattern {
		for k, _ := range row {
			if k == len(row)-1 {
				continue
			}

			mirror := isColumnMirror(k, row)
			if mirror {
				patternMirrorSet[k] += 1
			}
		}
	}

	for mirrorColumnIndex, count := range patternMirrorSet {
		if count == rowCount-matchDiff {
			patternMirrors = append(patternMirrors, mirrorColumnIndex)
		}
	}
	return patternMirrors
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
