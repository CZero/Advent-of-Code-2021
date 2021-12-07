package main

import (
	"fmt"
	"strconv"
	"strings"
)

var fuelmap = make(map[int]int)

func main() {
	// lines, _ := readLines("example.txt")
	lines, _ := readLines("input.txt")
	locations, low, high := GetLocations(lines)
	sweetspot := CalcSweetspot(locations, low, high)
	fmt.Printf("The sweetspot is position %d, with %d fuel needed\n", sweetspot, fuelmap[sweetspot])
}

// CalcSweetspot calculates the place with the least amount of fuel needed.
func CalcSweetspot(locations []int, low, high int) int {
	var sweetspot int
	for loc := low; loc <= high; loc++ {
		var fuelneeded int
		for _, crab := range locations {
			fuelneeded += CalcDistance(loc, crab)
		}
		fuelmap[loc] = fuelneeded
		if fuelneeded < fuelmap[sweetspot] || loc == low { // Either the least fuel needed or the first position
			sweetspot = loc
		}
	}
	return sweetspot
}

// CalcDistance calculates the distance between two points, where each step costs one more fuel
// This is calculated by the formula: S = n/2(n+1) (sum of natural numbers formula)
// https://www.numerade.com/questions/the-sum-s-of-the-first-n-natural-numbers-123-ldots-n-is-given-by-the-formula-sfracn2n1-how-many-cons/
func CalcDistance(a, b int) int {
	var distance int
	switch {
	case a > b:
		distance = a - b
	case b > a:
		distance = b - a
	default:
		distance = 0
	}
	answer := int(float64(distance) / 2 * (float64(distance) + 1))
	return answer
}

// GetLocations get the locations of the input
func GetLocations(input []string) (locations []int, low, high int) {
	fields := strings.Split(input[0], ",")
	for _, loc := range fields {
		l := SilentAtoi(loc)
		locations = append(locations, l)
		switch {
		case l > high:
			high = l
		case l < low:
			low = l
		}
	}
	return locations, low, high
}

// SilentAtoi converts string to int without returning errors
func SilentAtoi(input string) int {
	val, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return val
}
