package main

import (
	"fmt"
	"strconv"
)

func main() {
	// lines, _ := readLines("example.txt") // contains the example data
	lines, _ := readLines("input.txt") // contains the real data
	gamma := FindGamma(lines)
	gammanumber := ToBinary(gamma)
	epsilon := FindEpsilon(gamma)
	epsilonnumber := ToBinary(epsilon)
	fmt.Printf("Gamma:             %s = %d\nEpsilon:           %10s = %d\nPower consumption: %d * %d = %d\n\n", gamma, gammanumber, epsilon, epsilonnumber, gammanumber, epsilonnumber, gammanumber*epsilonnumber)
	oxygen := findOandO2(gamma, lines, true)
	oxygennumber := ToBinary(oxygen)
	co2 := findOandO2(epsilon, lines, false)
	co2number := ToBinary(co2)
	fmt.Printf("Oxygen generator:    %s = %d\nCO2 scrubber rating: %s = %d\nLife support rating: %d * %d = %d\n\n", oxygen, oxygennumber, co2, co2number, oxygennumber, co2number, oxygennumber*co2number)
}

// FindGamma finds the most common bits.
func FindGamma(input []string) (gamma string) {
	length := len(input[0])   // Length of the binaries
	zero := make(map[int]int) // Map of zero counts on the positions
	one := make(map[int]int)  // Map of one counts on the positions
	for _, line := range input {
		for pos, char := range line {
			switch string(char) {
			case "0":
				zero[pos]++
			case "1":
				one[pos]++
			default:
				panic("Invalid input: " + string(char))
			}
		}
	}
	for i := 0; i < length; i++ {
		if zero[i] > one[i] {
			gamma += "0"
		} else {
			gamma += "1"
		}
	}
	return gamma
}

// FindEpsilon finds the least common bits by inverting gamma
func FindEpsilon(gamma string) (epsilon string) {
	for _, char := range gamma {
		if string(char) == "0" {
			epsilon += "1"
		} else {
			epsilon += "0"
		}
	}
	return epsilon
}

// ToBinary takes a string binary number and returns it as an int
func ToBinary(input string) int {
	i, err := strconv.ParseInt(input, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

// findOandO2 takes the initial filter and the input. If it is looking for Oxygen it wil look for the most common (gamma),
// if looking for the CO2 Scrubber value, then for the least common (epsilon).
func findOandO2(filter string, input []string, oxygen bool) (found string) {
	heap := input
	var stack []string
	for i := 0; i < len(filter); i++ {
		for _, line := range heap {
			// We need to find the new filter (the new common or least in the new heap)
			if i > 0 {
				if oxygen {
					filter = FindGamma(heap)
				} else {
					filter = FindEpsilon(FindGamma(heap))
				}
			}
			if string(line[i]) == string(filter[i]) {
				stack = append(stack, line)
			}
		}
		heap = stack
		stack = nil
		if len(heap) == 1 {
			return heap[0]
		}
	}
	return ""
}
