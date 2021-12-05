package main

import (
	"fmt"
	"strconv"
	"strings"
)

type coords struct {
	x int // Horizontal
	y int // Vertical
}

var Grid = make(map[coords]int, 1000000) // The grid -> Optimize by first scanning for size if wanted, now I just used 1000x1000. Growing is expensive!

func main() {
	// lines, _ := readLines("example.txt")
	lines, _ := readLines("input.txt")
	for _, line := range lines {
		ResolveLine(line)
	}
	fmt.Printf("At %d points there is more than 1 cloud", FindCloudy())
}

// ResolveLine resolves a line (or a dot)
func ResolveLine(input string) {
	var (
		start coords // The start of the line
		stop  coords // The end of the line
	)
	line := strings.Split(input, "->") // [0] start, [1] end
	start = GetCoords(strings.TrimSpace(line[0]))
	stop = GetCoords(strings.TrimSpace(line[1]))
	switch {
	case start.x == stop.x && start.y != stop.y: // Vertical line
		lowest, highest := LowToHigh(start.y, stop.y)
		for i := lowest; i <= highest; i++ {
			Grid[coords{start.x, i}]++
		}
	case start.y == stop.y && start.x != stop.x: // Horizontal line
		lowest, highest := LowToHigh(start.x, stop.x)
		for i := lowest; i <= highest; i++ {
			Grid[coords{i, start.y}]++
		}
	case start.x != stop.x && start.y != stop.y: // Diagonal line, we are promised only 45 degrees!
		XRange := GetRange(start.x, stop.x) // Get all the x's lined up in a list
		YRange := GetRange(start.y, stop.y) // Get all the y's lined up in a list
		if len(XRange) != len(YRange) {
			panic("We were promised diagonals, they lied to us! Hissy fit!")
		}
		for i := 0; i < len(XRange); i++ { // Traverse the combined x,y steps
			Grid[coords{XRange[i], YRange[i]}]++
		}

	default: // If all are the same, it's a dot.
		Grid[coords{start.x, start.y}]++
	}
}

// GetCoords takes a string x,y and returns a coord.
func GetCoords(input string) coords {
	pos := strings.Split(input, ",")
	return coords{SilentAtoi(pos[0]), SilentAtoi(pos[1])}
}

// SilentAtoi just returns the int, not the err
func SilentAtoi(input string) int {
	res, err := strconv.Atoi(input)
	if err != nil {
		panic(err) // Ruh roh..
	}
	return res
}

// LowToHigh gives back 2 numbers in the order of low, high
func LowToHigh(a, b int) (low, high int) {
	if a > b {
		return b, a
	}
	return a, b // If we got here, it's the other option
}

// FindCloud loops through the set entries in the grid, to see which have more than 1 cloud
// It doesn't need to check all the grid entries, just the ones we've set.
func FindCloudy() int {
	var cloudsfound int
	for _, val := range Grid {
		if val > 1 {
			cloudsfound++
		}
	}
	return cloudsfound
}

// GetRange returns an []int list of all the steps from a -> b
func GetRange(a, b int) []int {
	var steps []int
	if a > b {
		for i := 0; i <= a-b; i++ {
			steps = append(steps, a-i)
		}
	} else {
		for i := 0; i <= b-a; i++ {
			steps = append(steps, a+i)
		}
	}
	return steps
}
