package main

import (
	"11a/uls"
	"fmt"
	"strconv"
)

type coord struct {
	r int // Row
	c int // Column
}

func main() {
	var (
		octopi     = make(map[coord]int) // The grid of octopi
		maxr       int                   // Max rows (grid height)
		maxc       int                   // Max colums (grid width)
		step       int                   // Step keeps track of the number of steps taken
		steps      = 100                 // The number of steps to calculate
		flashcount int                   // Flashcount keeps track of the number of flashes.
	)
	// lines, _ := readLines("smallexample.txt")
	// lines, _ := readLines("example.txt")
	lines, _ := readLines("input.txt")
	_ = lines
	maxr, maxc = GetOctopi(lines, octopi)
	PresentGrid(octopi, maxr, maxc, step, flashcount)
	for i := 0; i < steps; i++ {
		StepUp(octopi, maxr, maxc, &step, &flashcount)
		ResetFlashed(octopi, maxr, maxc)
		PresentGrid(octopi, maxr, maxc, step, flashcount)
	}
}

// ResetFlashed resets all flashed octopi to 0 after a round, and checks if all were flashed.
func ResetFlashed(octopi map[coord]int, maxr, maxc int) {
	for r := 0; r <= maxr; r++ {
		for c := 0; c <= maxc; c++ {
			if octopi[coord{r, c}] > 9 {
				octopi[coord{r, c}] = 0
			}
		}
	}
}

// StepUp runs through a step of increasing the values of the octopi and flashes them when they pass 9
func StepUp(octopi map[coord]int, maxr, maxc int, step, flashcount *int) {
	*step++
	for r := 0; r <= maxr; r++ {
		for c := 0; c <= maxc; c++ {
			octopi[coord{r, c}]++
			if octopi[coord{r, c}] == 10 {
				DoFlash(octopi, coord{r, c}, maxr, maxc, step, flashcount)
			}
		}
	}
}

// DoFlash executes a flash on an octopus. If it sets of another octopus it will call that flash too.
func DoFlash(octopi map[coord]int, Coord coord, maxr, maxc int, step, flashcount *int) {
	*flashcount++
	for rmod := -1; rmod < 2; rmod++ {
		for cmod := -1; cmod < 2; cmod++ {
			if Coord.r+rmod >= 0 && Coord.r+rmod <= maxr && Coord.c+cmod >= 0 && Coord.c+cmod <= maxc {
				if !(cmod == 0 && rmod == 0) {
					octopi[coord{Coord.r + rmod, Coord.c + cmod}]++
					if octopi[coord{Coord.r + rmod, Coord.c + cmod}] == 10 {
						DoFlash(octopi, coord{Coord.r + rmod, Coord.c + cmod}, maxr, maxc, step, flashcount)
					}
				}
			}
		}
	}
}

// GetOctopi reads the input of the puzzel into the Octopi grid
func GetOctopi(input []string, octopi map[coord]int) (maxr, maxc int) {
	maxr = len(input) - 1    // As the grid will go from 0 - x, we allready adjust the max r to 0 index.
	maxc = len(input[0]) - 1 // Ditto
	for r, line := range input {
		for c, octopus := range line {
			octopi[coord{r, c}] = uls.SilentAtoi(string(octopus))
		}
	}
	return maxr, maxc
}

// PresentGrid is a display tool to display the octopi spaced evenly.
func PresentGrid(octopi map[coord]int, maxr, maxc, step, flashcount int) {
	fmt.Printf("Step: %d\n\n", step)
	for r := 0; r <= maxr; r++ {
		var row string // The row we will present gets build
		for c := 0; c <= maxc; c++ {
			val := octopi[coord{r, c}]
			switch {
			case val < 10:
				row += strconv.Itoa(val) + "  "
			default:
				row += strconv.Itoa(val) + " "
			}
		}
		fmt.Println(row)
	}
	fmt.Printf("\nFlashcount: %d\n\n====================\n", flashcount)
}
