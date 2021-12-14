package main

import (
	"fmt"
	"sort"
)

func main() {
	// lines, _ := readLines("example.txt")
	lines, _ := readLines("input.txt")

	var results []int
	for _, line := range lines {
		results = append(results, CheckLine(line))
	}
	fmt.Printf("Gevonden: %#v\nThe middle value: %d\n", results, FindMiddleScore(results))
}

func CheckLine(line string) int {
	var openers []string
	// Corruption checker. Returns 0 on error.
	for _, char := range line {
		switch string(char) {
		case "(", "[", "{", "<":
			openers = append(openers, string(char))
		case ")":
			if openers[len(openers)-1] == "(" {
				openers = openers[:len(openers)-1]
			} else {
				return 0
			}
		case "]":
			if openers[len(openers)-1] == "[" {
				openers = openers[:len(openers)-1]
			} else {
				return 0
			}
		case "}":
			if openers[len(openers)-1] == "{" {
				openers = openers[:len(openers)-1]
			} else {
				return 0
			}
		case ">":
			if openers[len(openers)-1] == "<" {
				openers = openers[:len(openers)-1]
			} else {
				return 0
			}
		default:
			panic("illegal char" + string(char))
		}
	}
	// We survived the corruption checker. See if we have remainders. If so, score them.
	if len(openers) == 0 {
		return 0
	}
	return FinishLine(openers)
}

func FinishLine(openers []string) int {
	var score int
	for i := 0; i < len(openers); i++ {
		switch openers[len(openers)-1-i] {
		case "(":
			score = score*5 + 1
		case "[":
			score = score*5 + 2
		case "{":
			score = score*5 + 3
		case "<":
			score = score*5 + 4
		}
	}
	return score
}

func SumScores(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func FindMiddleScore(numbers []int) int {
	var nozeroes []int
	for _, num := range numbers {
		if num != 0 {
			nozeroes = append(nozeroes, num)
		}
	}
	sort.Ints(nozeroes)
	return (nozeroes[len(nozeroes)/2])
}
