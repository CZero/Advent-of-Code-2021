package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	newfish  = 8 // 9-1
	normfish = 6 // 7-1
	duration = 80
)

func main() {
	// lines, _ := readLines("example.txt")
	lines, _ := readLines("input.txt")
	school := FillSchool(lines)
	school = GrowSchool(school, duration)
	fmt.Println(CountFish(school))
}

func GrowSchool(school []int, duration int) []int {
	for d := 0; d < duration; d++ {
		var extrafish int
		for n, fish := range school {
			if fish == 0 {
				extrafish++
				school[n] = normfish
			} else {
				school[n]--
			}
		}
		for n := 1; n <= extrafish; n++ {
			school = append(school, newfish)
		}
		// fmt.Printf("After %d days: %v\n", d+1, school)
	}
	return school
}

func CountFish(school []int) int {
	return len(school)
}

func FillSchool(input []string) (school []int) {
	step := strings.Split(input[0], ",")
	for _, num := range step {
		school = append(school, SilentAtoi(num))
	}
	fmt.Println(school)
	return school
}

func SilentAtoi(a string) int {
	val, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return val
}
