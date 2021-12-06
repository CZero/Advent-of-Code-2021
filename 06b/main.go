package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	newfish  = 8 // 9-1
	normfish = 6 // 7-1
	duration = 256
)

func main() {
	// lines, _ := readLines("example.txt")
	lines, _ := readLines("input.txt")
	school := FillSchool(lines)
	fmt.Println(school)
	preschool := GrowSchool(school, duration)
	fmt.Println(CountFish(school), CountFish(preschool), CountFish(school)+CountFish(preschool))
}

// GrowSchool plays a day and sees what it does to the schools. Fishes produce new fish and become 6 again,
// the new fish first have to attend preschool
func GrowSchool(school map[int]int, duration int) map[int]int {
	// Since there are 2 types of fish. Newborn and normal. We first deal with normal
	preschool := make(map[int]int) // Prepare preschool for newlyborns!
	var graduates int              // Fishies in the proces of leaving their class
	var prev int                   // Fishes in the proces of coming to the class
	for d := 0; d < duration; d++ {
		// School on a brand new day.
		prev = 0
		for age := 6; age >= 0; age-- {
			graduates = school[age]
			school[age] = prev
			prev = graduates
		}
		school[6] = prev // Fishes which produced go back to 6

		// Nu de preschoolers (nieuwe vissen hebben 2 jaar meer nodig voor ze "6" zijn. Dat zijn er even
		// veel als dat er naar klas 6 gingen)
		for age := 1; age >= 0; age-- {
			graduates = preschool[age]
			preschool[age] = prev
			prev = graduates
		}
		school[6] += prev // Vissen die uit jaar het laatste jaar komen stromen in bij 6
		fmt.Println("Day ", d+1, " School: ", school, " Preschool: ", preschool)
	}
	return preschool
}

// CountFish counts the fish in a school
func CountFish(school map[int]int) (sum int) {
	for _, n := range school {
		sum += n
	}
	return sum
}

// FillSchool leest de eerste lichting uit
func FillSchool(input []string) (school map[int]int) {
	school = make(map[int]int)
	step := strings.Split(input[0], ",")
	for _, num := range step {
		school[SilentAtoi(num)]++
	}
	return school
}

// SilentAtoi returns no errors, only int.
func SilentAtoi(a string) int {
	val, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return val
}
