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
