package main

import (
	"fmt"
	"strconv"
)

func main() {
	lines, _ := readLines("input.txt")
	fmt.Printf("Times increased: %d\n", countIncrease(lines))
}

func countIncrease(input []string) int {
	var (
		prev           int // De vorige meting, om te onthouden
		increasedcount int // Aantal keren dat het meer was
	)
	for i, line := range input {
		depth, _ := strconv.Atoi(line)
		if i > 0 {
			if depth > prev {
				increasedcount++
			}
		}
		prev = depth
	}
	return increasedcount
}
