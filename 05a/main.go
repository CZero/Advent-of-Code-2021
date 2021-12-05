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

// ResolveLine resolves a line, but only of it's a dot or a horizontal or vertical line.
// Diagonal is bluntly ignored.
func ResolveLine(input string) {
	var (
		start coords // The start of the line
		stop  coords // The end of the line
	)
	line := strings.Split(input, "->") // 0 from, 1 to
	start = GetCoords(Trim(line[0]))
	stop = GetCoords(Trim(line[1]))
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
	case start.x != stop.x && start.y != stop.y: // Diagonal line
		// Ignore these
	default: // If all are the same, it's a dot.
		Grid[coords{start.x, start.y}]++
	}
}

// Trim = Quicktrim, easier to read code
func Trim(input string) string {
	return strings.TrimSpace(input)
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
		panic(err)
	}
	return res
}

// LowToHigh gives back 2 numbers in the order of low, high
func LowToHigh(a, b int) (low, high int) {
	if a > b {
		return b, a
	}
	return a, b
}

func FindCloudy() int {
	var cloudsfound int
	for _, val := range Grid {
		if val > 1 {
			cloudsfound++
		}
	}
	return cloudsfound
}
