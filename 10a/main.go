package main

import "fmt"

func main() {
	// lines, _ := readLines("example.txt")
	lines, _ := readLines("input.txt")

	var results []int
	for _, line := range lines {
		results = append(results, CheckLine(line))
	}
	fmt.Printf("Gevonden: %#v, sum = %d\n", results, SumScores(results))
}

func CheckLine(line string) int {
	var errorvalues = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	var openers []string
	for _, char := range line {
		switch string(char) {
		case "(", "[", "{", "<":
			openers = append(openers, string(char))
		case ")":
			if openers[len(openers)-1] == "(" {
				openers = openers[:len(openers)-1]
			} else {
				fmt.Printf("%s found, but last was: %s. So %d points.\n", string(char), openers[len(openers)-1], errorvalues[string(char)])
				return errorvalues[string(char)]
			}
		case "]":
			if openers[len(openers)-1] == "[" {
				openers = openers[:len(openers)-1]
			} else {
				fmt.Printf("%s found, but last was: %s. So %d points.\n", string(char), openers[len(openers)-1], errorvalues[string(char)])
				return errorvalues[string(char)]
			}
		case "}":
			if openers[len(openers)-1] == "{" {
				openers = openers[:len(openers)-1]
			} else {
				fmt.Printf("%s found, but last was: %s. So %d points.\n", string(char), openers[len(openers)-1], errorvalues[string(char)])
				return errorvalues[string(char)]
			}
		case ">":
			if openers[len(openers)-1] == "<" {
				openers = openers[:len(openers)-1]
			} else {
				fmt.Printf("%s found, but last was: %s. So %d points.\n", string(char), openers[len(openers)-1], errorvalues[string(char)])
				return errorvalues[string(char)]
			}
		default:
			panic("illegal char" + string(char))
		}
	}
	return 0
}

func SumScores(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return sum
}
