package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strconv"
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
	times, distances := parseInput(input)

	result := 1
	for race := 0; race < len(times); race++ {
		raceTime := times[race]
		raceDistance := distances[race]
		numberOfWays := getNumberOfWaysToWin(raceTime, raceDistance)
		result = result * numberOfWays
	}

	return result
}

func part2(input string) int {
	times, distances := parseInput(input)
	stringTimes := util.ConvertToStringSlice(times)
	bigTime, _ := strconv.Atoi(strings.Join(stringTimes, ""))
	stringDistances := util.ConvertToStringSlice(distances)
	bigDistance, _ := strconv.Atoi(strings.Join(stringDistances, ""))

	numberOfWays := getNumberOfWaysToWin(bigTime, bigDistance)
	return numberOfWays
}

func parseInput(input string) (times []int, distances []int) {

	rows := strings.Split(input, "\n")

	timeString := strings.Split(rows[0], ":")[1]
	distanceString := strings.Split(rows[1], ":")[1]

	timesSlice := strings.Split(timeString, " ")
	distanceSlice := strings.Split(distanceString, " ")

	times = util.ConvertToIntSlice(timesSlice)
	distances = util.ConvertToIntSlice((distanceSlice))

	return times, distances
}

func getNumberOfWaysToWin(raceTime int, raceDistance int) (numberOfWays int) {
	positive, negative := QuadraticFormula(1, -float64(raceTime), float64(raceDistance+1))
	numberOfWays = int(math.Floor(positive)) - int(math.Ceil(negative)) + 1

	return numberOfWays
}

func QuadraticFormula(a, b, c float64) (positive float64, negative float64) {
	square := b*b - (4 * a * c)
	squareRoot := math.Sqrt(square)
	positive = (-b + squareRoot) / (2 * a)
	negative = (-b - squareRoot) / (2 * a)
	return math.Max(positive, negative), math.Min(positive, negative)
}
