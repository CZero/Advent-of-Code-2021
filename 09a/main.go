package main

import (
	"fmt"
	"strconv"
)

type pos struct {
	r int // Row
	c int // Col
}

var caves = make(map[pos]int) // Caves height map
var lowpoints []pos

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
	fmt.Printf("Sum of the lowpoints = %d", sumLowpointsrisk(lowpoints))
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
