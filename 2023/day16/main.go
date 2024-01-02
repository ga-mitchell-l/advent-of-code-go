package main

import (
	_ "embed"
	"flag"
	"fmt"
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
	_ = parsed

	rowCount := len(parsed)
	columnCount := len(parsed[0])

	beam := Beam{
		row:       0,
		column:    0,
		direction: "E",
	}
	startingBeams := make([]Beam, 0)
	startingBeams = append(startingBeams, beam)

	energisedTiles := make(map[string]bool)
	floop := 0
	for len(startingBeams) > 0 {
		fmt.Println("-------")
		fmt.Println("starting beams", startingBeams)
		currentBeam := startingBeams[0]
		startingBeams = startingBeams[1:]

		stop := false

		for currentBeam.row >= 0 && currentBeam.row < rowCount &&
			currentBeam.column >= 0 && currentBeam.column < columnCount && !stop && floop < 10 {
			floop++
			currentTile := parsed[currentBeam.row][currentBeam.column]
			tileIndex := strings.Join([]string{strconv.Itoa(currentBeam.row), strconv.Itoa(currentBeam.column)}, ",")
			energisedTiles[tileIndex] = true

			fmt.Println("- - - ")
			fmt.Println("tile index:", tileIndex)
			fmt.Println("current tile:", currentTile)
			fmt.Println("current direction", currentBeam.direction)

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
					fmt.Println("east split")
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
					currentBeam.column++
					nextDirection = "S"
				case downMirror:
					currentBeam.column--
					nextDirection = "N"
				case vertSplit:
					stop = true
					startingBeams = splitBeam(previousBeam, startingBeams, []string{"N", "S"})
				}
			}

			currentBeam.direction = nextDirection

		}

		if !stop {
			fmt.Println("reached edge")
		}
	}

	fmt.Println(energisedTiles)

	return len(energisedTiles)
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
	return 0
}

func parseInput(input string) (ans [][]string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, strings.Split(line, ""))
	}
	return ans
}
