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

type Lense struct {
	focalLength int
	label       string
}

func part2(input string) int {
	parsed := parseInput(input)
	boxes := getBoxes(parsed)

	fmt.Println(boxes)
	return 0
}

func getBoxes(parsed []string) map[int][]Lense {
	boxes := make(map[int][]Lense)

	for _, value := range parsed {
		equals := strings.Split(value, "=")
		dash := strings.Split(value, "-")

		if len(dash) > 1 {
			box, newBoxContents := getDashBoxContents(dash, boxes)
			boxes[box] = newBoxContents
		}
		if len(equals) > 1 {
			box, newBoxContents := getEqualsBoxContents(equals, boxes)
			boxes[box] = newBoxContents
		}
	}
	return boxes
}

func getEqualsBoxContents(equals []string, boxes map[int][]Lense) (int, []Lense) {
	label := equals[0]
	focalLength, _ := strconv.Atoi(equals[1])

	lense := Lense{
		focalLength,
		label,
	}

	box := getInitialisationStep(label)
	currentBoxContents := boxes[box]

	filter := func(s Lense) bool { return s.label == label }
	matchingLensesInBox := util.Filter(currentBoxContents, filter)
	if len(matchingLensesInBox) == 0 {
		currentBoxContents = append(currentBoxContents, lense)
	} else {
		for index, val := range currentBoxContents {
			if val.label == label {
				val.focalLength = focalLength
			}
			currentBoxContents[index] = val
		}
	}
	return box, currentBoxContents
}

func getDashBoxContents(dash []string, boxes map[int][]Lense) (int, []Lense) {
	label := dash[0]
	box := getInitialisationStep(label)

	currentBoxContents := boxes[box]
	newBoxContents := currentBoxContents
	for index, value := range currentBoxContents {
		_ = index
		_ = value

		if value.label == label {
			newBoxContents = append(currentBoxContents[:index], currentBoxContents[index+1:]...)
		}
	}
	return box, newBoxContents
}

func parseInput(input string) (ans []string) {
	return strings.Split(input, ",")
}
