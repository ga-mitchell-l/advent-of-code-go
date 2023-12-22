package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"

	pcre "github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"

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

func findRepeatedSequences(series string) string {
	re, _ := pcre.Compile(`(.*)\1+`, 0)
	matches := re.MatcherString(series, 0).GroupString(1)
	return matches
}

func part2(input string) int {

	parsed := parseInput(input)
	cycles := 10000
	cycleLoads := make([]string, 0)

	for i := 0; i < cycles; i++ {
		parsed = cycleRocks(parsed)
		load := calculateLoad(parsed)
		cycleLoads = append(cycleLoads, strconv.Itoa(load))
	}

	match := false
	matchCycleLoads := cycleLoads
	matchString := ""
	var matchStringSplit []int
	skipped := -1
	for match != true && len(matchCycleLoads) > 0 {
		skipped++
		cycleLoadsString := strings.Join(matchCycleLoads, ",")
		matchString = findRepeatedSequences(cycleLoadsString)

		matchCycleLoads = matchCycleLoads[1:]
		matchStringSplit = util.ConvertToIntSlice(strings.Split(matchString, ","))
		match = len(matchStringSplit) > 1
	}

	var previousMatchStringSplit []int

	for len(matchStringSplit) > 1 {
		previousMatchStringSplit = matchStringSplit
		matchString = findRepeatedSequences(matchString)
		matchStringSplit = util.ConvertToIntSlice(strings.Split(matchString, ","))
	}

	repeatCycle := len(previousMatchStringSplit)
	_ = repeatCycle
	resultCycles := 1000000000

	resultIndex := (resultCycles - 1 - skipped) % repeatCycle
	result := previousMatchStringSplit[resultIndex]

	return result
}

func cycleRocks(parsed [][]string) [][]string {
	northRocks := moveRocks(parsed, 1)
	transposedNorthRocks := util.Transpose(northRocks)
	transposedWestRocks := moveRocks(transposedNorthRocks, 1)
	westRocks := util.Transpose(transposedWestRocks)
	southRocks := moveRocks(westRocks, -1)
	transposedSouthRocks := util.Transpose(southRocks)
	transposedEastRocks := moveRocks(transposedSouthRocks, -1)
	eastRocks := util.Transpose(transposedEastRocks)
	return eastRocks
}

func parseInput(input string) (ans [][]string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, strings.Split(line, ""))
	}
	return ans
}
