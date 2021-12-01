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
		lastThree      []int // Contains the last three depths
		depth          int   // The depths this round
		previousSum    int   // The sum of the last three depths
		increasedCount int   // Time the depth was more than the previous 3
	)
	for _, line := range input {
		depth, _ = strconv.Atoi(line)
		if len(lastThree) == 3 {
			lastThree = lastThree[1:]
		}
		lastThree = append(lastThree, depth)
		if len(lastThree) == 3 {
			sum := 0
			for _, n := range lastThree {
				sum += n
			}
			if previousSum > 0 {
				if sum > previousSum {
					increasedCount++
				}
			}
			previousSum = sum
		}
	}
	return increasedCount
}
