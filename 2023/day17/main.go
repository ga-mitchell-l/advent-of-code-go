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
	row    int
	column int
	value  int
	dist   float64
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
	_ = parsed

	// Because it is difficult to keep the top-heavy crucible going in a straight line for very long,
	// it can move at most three blocks in a single direction before it must turn 90 degrees left or right.
	// The crucible also can't reverse direction;
	//after entering each city block, it may only turn left, continue straight, or turn right.

	maxBlockMoves := 3
	_ = maxBlockMoves
	maxRows := len(parsed)
	maxCols := len(parsed[0])
	dist, queue := getDistQueue(parsed)

	neighbourPositions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for len(queue) > 0 {
		// first run through, this block is the source block
		minDistBlock := slices.MinFunc[[]Block, Block](queue, minDistBlockFunc)
		fmt.Println("min:", minDistBlock)

		// remove current block from queue
		minQueueIndex := GetBlockSliceIndex(queue, minDistBlock.row, minDistBlock.column)
		queue = append(queue[:minQueueIndex], queue[minQueueIndex+1:]...)

		// for each neighbour block
		for _, position := range neighbourPositions {
			neighbourRow := minDistBlock.row + position[0]
			neighbourCol := minDistBlock.column + position[1]
			neighbourDistIndex := GetBlockSliceIndex(dist, neighbourRow, neighbourCol)
			if neighbourDistIndex == -1 {
				// check neighbour exists
				continue
			}

			neighbour := dist[neighbourDistIndex]
			alt := minDistBlock.dist + float64(neighbour.value)

			if alt < neighbour.dist {
				neighbourQueueIndex := GetBlockSliceIndex(queue, neighbourRow, neighbourCol)
				if neighbourQueueIndex > -1 {
					queue[neighbourQueueIndex].dist = alt
				}
				dist[neighbourDistIndex].dist = alt
			}
		}

	}

	fmt.Println(foo)
	fmt.Println(dist)
	fmt.Println(queue)

	endPointIndex := GetBlockSliceIndex(dist, maxRows-1, maxCols-1)
	endPoint := dist[endPointIndex]

	return int(endPoint.dist)
}

func getDistQueue(parsed [][]string) ([]Block, []Block) {
	dist := make([]Block, 0)
	queue := make([]Block, 0)

	for rowIndex, row := range parsed {
		for columnIndex, blockValue := range row {
			distValue := math.Inf(1)
			intBlockVal, _ := strconv.Atoi(blockValue)
			if rowIndex == 0 && columnIndex == 0 {
				distValue = 0
			}
			block := Block{
				row:    rowIndex,
				column: columnIndex,
				dist:   distValue,
				value:  intBlockVal,
			}
			dist = append(dist, block)
			queue = append(queue, block)
		}
	}
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
