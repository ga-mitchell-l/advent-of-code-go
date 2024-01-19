package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"slices"
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

type Block struct {
	row     int
	column  int
	value   int
	dist    float64
	history []string
}

type Position struct {
	rowDiff   int
	colDiff   int
	direction string
}

var minDistBlockFunc = func(A, B Block) int {
	if A.dist < B.dist {
		return -1
	} else if A.dist > B.dist {
		return 1
	} else {
		return 0
	}
}

func GetBlockSliceIndex(queue []Block, row int, column int) (ret int) {
	for index, value := range queue {
		if row == value.row && column == value.column {
			return index
		}
	}
	return -1
}

func part1(input string) int {
	parsed := parseInput(input)

	// Because it is difficult to keep the top-heavy crucible going in a straight line for very long,
	// it can move at most three blocks in a single direction before it must turn 90 degrees left or right.
	// The crucible also can't reverse direction;
	//after entering each city block, it may only turn left, continue straight, or turn right.

	maxRows := len(parsed)
	maxCols := len(parsed[0])
	dist, queue := getDistQueue(parsed)

	neighbourPositions := []Position{
		{rowDiff: -1, colDiff: 0, direction: "N"},
		{rowDiff: 0, colDiff: 1, direction: "E"},
		{rowDiff: 1, colDiff: 0, direction: "S"},
		{rowDiff: 0, colDiff: -1, direction: "W"},
	}

	debugRow := 0
	debugColumn := 2

	floop := 0
	for len(queue) > 0 && floop < 1000 {
		floop++
		// first run through, this block is the source block
		minDistBlock := slices.MinFunc[[]Block, Block](queue, minDistBlockFunc)

		if debugRow == minDistBlock.row && debugColumn == minDistBlock.column {
			fmt.Println("------")
			fmt.Println("row:", minDistBlock.row, "col:", minDistBlock.column, "val:", minDistBlock.value, "dist:", minDistBlock.dist, "hist:", minDistBlock.history)
			fmt.Println("- - - - ")
		}

		// remove current block from queue
		minQueueIndex := GetBlockSliceIndex(queue, minDistBlock.row, minDistBlock.column)
		queue = append(queue[:minQueueIndex], queue[minQueueIndex+1:]...)
		history := minDistBlock.history
		skipDirection := ""
		if len(history) > 2 {
			i := history[len(history)-1]
			j := history[len(history)-2]
			k := history[len(history)-3]

			if i == j && j == k {
				skipDirection = i
			}
		}

		// for each neighbour block
		for _, position := range neighbourPositions {
			if debugRow == minDistBlock.row && debugColumn == minDistBlock.column {
				fmt.Println("position:", position)
			}

			// check we are not reversing direction
			if position.direction == getReverseDirection(history) {
				if debugRow == minDistBlock.row && debugColumn == minDistBlock.column {
					fmt.Println("reverse direction :C")
					fmt.Println("- - - -")
				}
				continue
			}

			// check we haven't been going in the same direction for too long
			if position.direction == skipDirection {
				if debugRow == minDistBlock.row && debugColumn == minDistBlock.column {
					fmt.Println("too far same direction :C")
					fmt.Println("- - - -")
				}
				continue
			}

			// check neighbour exists
			neighbourRow := minDistBlock.row + position.rowDiff
			neighbourCol := minDistBlock.column + position.colDiff
			neighbourDistIndex := GetBlockSliceIndex(dist, neighbourRow, neighbourCol)
			if neighbourDistIndex == -1 {
				if debugRow == minDistBlock.row && debugColumn == minDistBlock.column {
					fmt.Println("does not exist :C")
					fmt.Println("- - - -")
				}
				continue
			}

			neighbourDist := dist[neighbourDistIndex]
			if debugRow == minDistBlock.row && debugColumn == minDistBlock.column {
				fmt.Println("neighbour:", neighbourDist)
			}
			alt := minDistBlock.dist + float64(neighbourDist.value)
			if debugRow == minDistBlock.row && debugColumn == minDistBlock.column {
				fmt.Println("alt:", alt)
			}

			if alt < neighbourDist.dist {
				if debugRow == minDistBlock.row && debugColumn == minDistBlock.column {
					fmt.Println("updating dist")
					fmt.Println("new history of neighbour:", append(history, position.direction))
				}
				neighbourQueueIndex := GetBlockSliceIndex(queue, neighbourRow, neighbourCol)
				if neighbourQueueIndex > -1 {
					queue[neighbourQueueIndex].dist = alt
					queue[neighbourQueueIndex].history = append(history, position.direction)
				}
				dist[neighbourDistIndex].dist = alt
				dist[neighbourDistIndex].history = append(history, position.direction)
			} else {
				if debugRow == minDistBlock.row && debugColumn == minDistBlock.column {
					fmt.Println("dist not changing")
				}
			}
			if debugRow == minDistBlock.row && debugColumn == minDistBlock.column {
				fmt.Println("- - - -")
			}
		}

	}

	endPointIndex := GetBlockSliceIndex(dist, maxRows-1, maxCols-1)
	endPoint := dist[endPointIndex]
	fmt.Println(endPoint.history)

	fmt.Println(dist[GetBlockSliceIndex(dist, 0, 0)])
	fmt.Println(dist[GetBlockSliceIndex(dist, 0, 1)])
	fmt.Println(dist[GetBlockSliceIndex(dist, 0, 2)])
	fmt.Println(dist[GetBlockSliceIndex(dist, 0, 3)])

	return int(endPoint.dist)
}

func getReverseDirection(s []string) string {
	if len(s) == 0 {
		return ""
	}
	switch s[len(s)-1] {
	case "N":
		return "S"
	case "E":
		return "W"
	case "S":
		return "N"
	case "W":
		return "E"
	}
	return ""
}

func getDistQueue(parsed [][]string) ([]Block, []Block) {
	dist := make([]Block, 0)
	queue := make([]Block, 0)

	for rowIndex, row := range parsed {
		for columnIndex, blockValue := range row {
			intBlockVal, _ := strconv.Atoi(blockValue)
			block := Block{
				row:     rowIndex,
				column:  columnIndex,
				dist:    math.Inf(1),
				value:   intBlockVal,
				history: make([]string, 0),
			}

			dist = append(dist, block)
			queue = append(queue, block)
		}
	}

	up := Block{
		row:     -1,
		column:  0,
		dist:    0,
		value:   0,
		history: make([]string, 0),
	}
	left := Block{
		row:     0,
		column:  -1,
		dist:    0,
		value:   0,
		history: make([]string, 0),
	}

	dist = append(dist, up, left)
	queue = append(queue, up, left)

	return dist, queue
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
