package main

import (
	"fmt"
	"strconv"
)

type pos struct {
	r int // Row
	c int // Col
}

var caves = make(map[pos]int)    // Caves height map
var visited = make(map[pos]bool) // Set if a point is checked for basin
var lowpoints []pos              // All the found lowpoints (starting points to basins)
var basins = make(map[pos][]pos) // A map of the basins

func main() {
	// lines, _ := readLines("example.txt")
	lines, _ := readLines("input.txt")
	ReadMap(lines)
	rowmax := len(lines)
	colmax := len(lines[0])
	for r := 0; r < rowmax; r++ {
		for c := 0; c < colmax; c++ {
			if CheckLowpoint(r, c, rowmax, colmax) {
				lowpoints = append(lowpoints, pos{r, c})
			}
		}
	}
	fmt.Println("Lowpoints: ", lowpoints)
	fmt.Printf("Sum of the lowpoints = %d\n\n", sumLowpointsrisk(lowpoints))
	for _, pos := range lowpoints {
		MapBasin(pos, rowmax, colmax, visited, pos)
	}
	fmt.Printf("The mapbasin knows these coords:\n%v\n", basins)
	fmt.Println(GetStats(basins))
}

func GetStats(basins map[pos][]pos) string {
	var (
		sizes   []int
		sum     int
		multi   = 1
		biggest [3]int
	)
	for _, basin := range basins {
		sizes = append(sizes, len(basin))
		sum += len(basin)
		multi *= len(basin)
		switch { // Hmm, deze moet beter kunnen?
		case len(basin) > biggest[0]:
			biggest[2] = biggest[1]
			biggest[1] = biggest[0]
			biggest[0] = len(basin)
		case len(basin) > biggest[1]:
			biggest[2] = biggest[1]
			biggest[1] = len(basin)
		case len(basin) > biggest[2]:
			biggest[2] = len(basin)
		}
	}
	top3multi := biggest[0] * biggest[1] * biggest[2]
	return fmt.Sprintf("There were %d basins found, with the follwing sizes:\n%v\n\nThis is a total of: %d (multiplied: %d)\nThe top 3 were: %v (multiplied: %d)", len(basins), sizes, sum, multi, biggest, top3multi)
}

func MapBasin(point pos, rowmax, colmax int, visited map[pos]bool, basin pos) {
	// fmt.Println(visited)
	if visited[point] { // We've been here allready
		return
	}
	visited[point] = true
	if caves[point] == 9 { // End of a basin
		return
	} else {
		basins[basin] = append(basins[basin], point)
		// basin = append(basin, point)
	}

	if point.r > 0 {
		MapBasin(pos{point.r - 1, point.c}, rowmax, colmax, visited, basin)
	}
	if point.r < rowmax-1 {
		MapBasin(pos{point.r + 1, point.c}, rowmax, colmax, visited, basin)
	}
	if point.c > 0 {
		MapBasin(pos{point.r, point.c - 1}, rowmax, colmax, visited, basin)
	}
	if point.c < colmax-1 {
		MapBasin(pos{point.r, point.c + 1}, rowmax, colmax, visited, basin)
	}
}

func CheckLowpoint(r, c, rowmax, colmax int) bool {
	lowpoint := true
	if r > 0 {
		if caves[pos{r, c}] >= caves[pos{r - 1, c}] {
			lowpoint = false
		}
	}
	if r < rowmax-1 {
		if caves[pos{r, c}] >= caves[pos{r + 1, c}] {
			lowpoint = false
		}
	}
	if c > 0 {
		if caves[pos{r, c}] >= caves[pos{r, c - 1}] {
			lowpoint = false
		}
	}
	if c < colmax-1 {
		if caves[pos{r, c}] >= caves[pos{r, c + 1}] {
			lowpoint = false
		}
	}
	return lowpoint
}

func sumLowpointsrisk(lowpoints []pos) (sum int) {
	for _, pos := range lowpoints {
		sum += 1 + caves[pos]
	}
	return sum
}

func ReadMap(input []string) {
	for row, line := range input {
		for col, height := range line {
			caves[pos{row, col}] = SilentAtoi(string(height))
		}
	}
}

func SilentAtoi(input string) int {
	val, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return val
}
