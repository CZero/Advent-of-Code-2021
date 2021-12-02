package main

import (
	"fmt"
	"strconv"
	"strings"
)

type loc struct { // The location of the submarine
	dep int // The depth (bigger is deeper!)
	hor int // The horizontal position
}

func main() {
	// lines, _ := readLines("example.txt") // lines contains the example input []string
	lines, _ := readLines("input.txt") // lines contains the challenge input []string
	_ = lines
	Loc := ResolveLocation(lines)
	fmt.Printf("The location of the sub is: %d,%d.\nMultiplied: %d", Loc.hor, Loc.dep, Loc.hor*Loc.dep)
}

func ResolveLocation(input []string) (Loc loc) {
	for _, line := range input {
		Loc = DoAction(line, Loc)
	}
	return Loc
}

func DoAction(line string, Loc loc) loc {
	instructions := strings.Fields(line)
	inc, err := strconv.Atoi(instructions[1])
	if err != nil {
		panic(err)
	}
	switch instructions[0] {
	case "forward":
		Loc.hor += inc
	case "down":
		Loc.dep += inc
	case "up":
		Loc.dep -= inc
	default:
		panic("Gotten a strange input!" + instructions[0])
	}
	return Loc
}
