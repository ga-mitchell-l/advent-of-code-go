package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"advent-of-code-go/util"
)

//go:embed input.txt
var input string

//go:embed example1.txt
var example1 string

var emptySpace, upMirror, downMirror, vertSplit, horizSplit = ".", "/", "\\", "|", "-"

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

type Beam struct {
	row       int
	column    int
	direction string
}

func part1(input string) int {
	parsed := parseInput(input)

	startingBeam := Beam{
		row:       0,
		column:    0,
		direction: "E",
	}

	return getEnergisedCount(parsed, startingBeam)
}

func getEnergisedCount(parsed [][]string, beam Beam) int {
	rowCount := len(parsed)
	columnCount := len(parsed[0])
	startingBeams := make([]Beam, 0)
	startingBeams = append(startingBeams, beam)

	energisedTiles := make(map[string]bool)
	previousBeams := make(map[string]bool)
	for len(startingBeams) > 0 {
		currentBeam := startingBeams[0]
		startingBeams = startingBeams[1:]
		beamIndex := strconv.Itoa(currentBeam.row) + strconv.Itoa(currentBeam.column) + currentBeam.direction

		stop := previousBeams[beamIndex]
		previousBeams[beamIndex] = true

		for currentBeam.row >= 0 && currentBeam.row < rowCount &&
			currentBeam.column >= 0 && currentBeam.column < columnCount && !stop {
			currentTile := parsed[currentBeam.row][currentBeam.column]
			tileIndex := strings.Join([]string{strconv.Itoa(currentBeam.row), strconv.Itoa(currentBeam.column)}, ",")
			energisedTiles[tileIndex] = true

			nextDirection := ""
			previousBeam := currentBeam
			switch currentBeam.direction {
			case "N":
				switch currentTile {
				case emptySpace, vertSplit:
					currentBeam.row--
					nextDirection = currentBeam.direction
				case upMirror:
					currentBeam.column++
					nextDirection = "E"
				case downMirror:
					currentBeam.column--
					nextDirection = "W"
				case horizSplit:
					stop = true
					startingBeams = splitBeam(previousBeam, startingBeams, []string{"E", "W"})
				}
			case "E":
				switch currentTile {
				case emptySpace, horizSplit:
					currentBeam.column++
					nextDirection = currentBeam.direction
				case upMirror:
					currentBeam.row--
					nextDirection = "N"
				case downMirror:
					currentBeam.row++
					nextDirection = "S"
				case vertSplit:
					stop = true
					startingBeams = splitBeam(previousBeam, startingBeams, []string{"N", "S"})
				}
			case "S":
				switch currentTile {
				case emptySpace, vertSplit:
					currentBeam.row++
					nextDirection = currentBeam.direction
				case upMirror:
					currentBeam.column--
					nextDirection = "W"
				case downMirror:
					currentBeam.column++
					nextDirection = "E"
				case horizSplit:
					stop = true
					startingBeams = splitBeam(previousBeam, startingBeams, []string{"E", "W"})
				}
			case "W":
				switch currentTile {
				case emptySpace, horizSplit:
					currentBeam.column--
					nextDirection = currentBeam.direction
				case upMirror:
					currentBeam.row++
					nextDirection = "S"
				case downMirror:
					currentBeam.row--
					nextDirection = "N"
				case vertSplit:
					stop = true
					startingBeams = splitBeam(previousBeam, startingBeams, []string{"N", "S"})
				}
			}

			currentBeam.direction = nextDirection

		}
	}

	energisedCount := len(energisedTiles)
	return energisedCount
}

func splitBeam(currentBeam Beam, startingBeams []Beam, directions []string) []Beam {
	newBeamE := Beam{
		row:       currentBeam.row,
		column:    currentBeam.column,
		direction: directions[0],
	}
	startingBeams = append(startingBeams, newBeamE)
	newBeamW := Beam{
		row:       currentBeam.row,
		column:    currentBeam.column,
		direction: directions[1],
	}
	startingBeams = append(startingBeams, newBeamW)
	return startingBeams
}

func part2(input string) int {
	parsed := parseInput(input)
	_ = parsed

	rowCount := len(parsed)
	columnCount := len(parsed[0])
	energisedCount := make([]int, 0)

	for i := 0; i < rowCount; i++ {
		leftBeam := Beam{
			row:       i,
			column:    0,
			direction: "E",
		}
		leftCount := getEnergisedCount(parsed, leftBeam)
		energisedCount = append(energisedCount, leftCount)

		rightBeam := Beam{
			row:       i,
			column:    columnCount - 1,
			direction: "W",
		}
		rightCount := getEnergisedCount(parsed, rightBeam)
		energisedCount = append(energisedCount, rightCount)
	}

	for i := 0; i < columnCount; i++ {
		topBeam := Beam{
			row:       0,
			column:    i,
			direction: "S",
		}
		topCount := getEnergisedCount(parsed, topBeam)
		energisedCount = append(energisedCount, topCount)

		bottomBeam := Beam{
			row:       rowCount - 1,
			column:    i,
			direction: "N",
		}
		bottomCount := getEnergisedCount(parsed, bottomBeam)
		energisedCount = append(energisedCount, bottomCount)

	}

	return slices.Max(energisedCount)
}

func parseInput(input string) (ans [][]string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, strings.Split(line, ""))
	}
	return ans
}
