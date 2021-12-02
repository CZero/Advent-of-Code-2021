package main

import (
	"fmt"
	"strconv"
	"strings"
)

type loc struct { // The location of the submarine
	dep int // The depth (bigger is deeper!)
	hor int // The horizontal position
	aim int // The aim amount
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
	case "forward": // Forward by x, depth is Aim*x
		Loc.hor += inc
		Loc.dep += inc * Loc.aim
	case "down": // Increases aim by x units.
		Loc.aim += inc
	case "up": // Increases aim by x units.
		Loc.aim -= inc
	default:
		panic("Gotten a strange input!" + instructions[0])
	}
	return Loc
}
