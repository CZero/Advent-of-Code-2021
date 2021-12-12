package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/** Display numbers

  0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg

ocurance streep alle cijfers:   dus:
a = 8                           4 = e
b = 6                           6 = b
c = 8                           7 = d,g
d = 7                           8 = a,f,c
e = 4
f = 8
g = 7

aantal strepen is getal:
2 = 1
3 = 7
4 = 4
5 = 2,3,5
6 = 0,6,9
7 = 8

voorkomen:

a = 0,  2,3,4,5,6,7,8,9
b = 0,      4,5,6,  8,9
c = 0,1,2,3,4,    7,8,9
d =     2,3,4,5,6,  8,9
e = 0,  2,      6,  8,
f = 0,1,  3,4,5,6,7,8,9
g = 0,  2,3,  5,6,  8,9

 **/

/**
Bij de uniques:
       1, 4, 7, 8
a, 2 =     , 7, 8
b, 2 =  , 4,  , 8
c, 4 = 1, 4, 7, 8
d, 2 =  , 4,  , 8
e, 1 =  ,  ,  , 8
f, 4 = 1, 4, 7, 8
g, 1 =  ,  ,  , 8
**/

func main() {
	// lines, _ := readLines("example.txt")
	lines, _ := readLines("input.txt")
	var totalfound int
	for _, line := range lines {
		totalfound += ResolveLine(line)
	}
	fmt.Println("Total: ", totalfound)
}

func ResolveLine(input string) int {
	var solution = make(map[string]string) // solution[segment]key
	var segmentoccurance = make(map[string]int)
	var uniques []string               // 1, 4, 7, 8
	parts := strings.Split(input, "|") // 0 is digits, 1 is display

	SegmentOccurance(strings.Fields(parts[0]), segmentoccurance, solution) // Now we now the number of occurance of segments, it get's us the first 2 segments
	uniques = GetUniques(parts[0])
	// fmt.Printf("Segment occurance: %v\n", segmentoccurance)
	// fmt.Printf("The solution thusfar: %v\n", solution)
	var uniquesegmentoccurance = make(map[string]int)
	SegmentOccurance(uniques, uniquesegmentoccurance, solution)
	// fmt.Printf("Unique segment occurance: %v\n", uniquesegmentoccurance)

	ResolveUniqueSegments(uniques, uniquesegmentoccurance, solution) // Nu hebben we A, B, D, E, G opgelost c en f alleen nog dus.
	// fmt.Printf("The solution thusfar (resolved unique segments): %v\n", solution)

	ResolveCandF(parts[0], solution)
	// fmt.Printf("The solution complete (resolved C and F): %v\n", solution)

	var numbersmap = make(map[string]string)
	MapNumbers(parts[0], solution, numbersmap)
	// fmt.Println("Numbersmap: ", numbersmap)
	var resolvednumber string
	for _, number := range strings.Fields(parts[1]) {
		tempsort := strings.Split(number, "")
		sort.Strings(tempsort)
		number = strings.Join(tempsort, "")
		// fmt.Println(number)
		resolvednumber += numbersmap[number]
	}
	fmt.Println("Yeah, the number = ", resolvednumber)
	// fmt.Println(segmentoccurance)
	return SilentAtoi(resolvednumber)
}

func MapNumbers(numbers string, solution map[string]string, numbersmap map[string]string) {
	var (
		zero = "ABCEFG"
		// one   = "CF"
		two   = "ACDEG"
		three = "ACDFG"
		// four  = "BCDF"
		five = "ABDFG"
		six  = "ABDEFG"
		// seven = "ACF"
		// eight = "ABCDEFG"
		nine = "ABCDFG"
	)
	for _, number := range strings.Fields(numbers) {
		tempsort := strings.Split(number, "")
		sort.Strings(tempsort)
		number = strings.Join(tempsort, "")

		switch len(number) {
		case 2:
			numbersmap[number] = "1"
		case 3:
			numbersmap[number] = "7"
		case 4:
			numbersmap[number] = "4"
		case 7:
			numbersmap[number] = "8"
		case 5: // 2,3,5
			// fmt.Printf("Het nummer dat we zoeken: %s\n", number)
			found := true
			for _, c := range number {
				if !IsIn(solution[string(c)], two) {
					found = false
				} else {
					// fmt.Printf("%s zit in %s\n", string(c), two)
				}
			}
			if found {
				numbersmap[number] = "2"
			}
			found = true
			for _, c := range number {
				if !IsIn(solution[string(c)], three) {
					found = false
				} else {
					// fmt.Printf("%s zit in %s\n", string(c), three)
				}
			}
			if found {
				numbersmap[number] = "3"
			}
			found = true
			for _, c := range number {
				if !IsIn(solution[string(c)], five) {
					found = false
				} else {
					// fmt.Printf("%s zit in %s\n", string(c), five)
				}
			}
			if found {
				numbersmap[number] = "5"
			}
		case 6: // 0, 6, 9
			found := true
			for _, c := range number {
				if !IsIn(solution[string(c)], zero) {
					found = false
				} else {
					// fmt.Printf("%s zit in %s\n", string(c), zero)
				}
			}
			if found {
				numbersmap[number] = "0"
			}
			found = true
			for _, c := range number {
				if !IsIn(solution[string(c)], six) {
					found = false
				} else {
					// fmt.Printf("%s zit in %s\n", string(c), six)
				}
			}
			if found {
				numbersmap[number] = "6"
			}
			found = true
			for _, c := range number {
				if !IsIn(solution[string(c)], nine) {
					found = false
				} else {
					// fmt.Printf("%s zit in %s\n", string(c), nine)
				}
			}
			if found {
				numbersmap[number] = "9"
			}
		}

	}
	return
}

func ResolveCandF(numbers string, solution map[string]string) {
	// Only two to resolve left: C and F. C zit niet in 5 en 6 waarbij die ook niet even lang zijn (5 en 6 lang toevallig)
	// F zit niet in 2, die is 5 lang.
	// Dus als een segment maar in 1 niet inzit is het F, anders is het C.
	for _, char := range "abcdefg" {
		if _, ok := solution[string(char)]; !ok {
			var count int
			for _, number := range strings.Fields(numbers) {
				if IsIn(string(char), number) {
					count++
				}
			}
			if count == 9 {
				solution[string(char)] = "F"
			} else {
				solution[string(char)] = "C"
			}
		}
	}
}

func ResolveUniqueSegments(uniques []string, uniquesegmentoccurance map[string]int, solution map[string]string) {
	for segment, occurance := range uniquesegmentoccurance {
		switch occurance {
		case 2: // A, B of D. A zit enkel in zeven en acht. B en D zitten in vier en 8, maar b hebben we al gevonden!
			switch {
			case IsIn(segment, uniques[2]) && IsIn(segment, uniques[3]):
				solution[segment] = "A"
			case IsIn(segment, uniques[1]) && IsIn(segment, uniques[3]):
				if solution[segment] != "B" {
					solution[segment] = "D"
				}
			}
		case 1: // E of G, maar e hebben we al gevonden!
			if solution[segment] != "E" {
				solution[segment] = "G"
			}
		case 4: // c of f. Deze weten we helaas nog niet, kan ik zo ook niet beredeneren.
		}
	}
}

func SegmentOccurance(input []string, segmentoccurance map[string]int, solution map[string]string) {
	for _, letter := range input {
		for _, v := range letter {
			segmentoccurance[string(v)]++
		}
	}
	if len(input) > 4 {
		for k, v := range segmentoccurance {
			if v == 4 {
				solution[k] = "E"
			}
			if v == 6 {
				solution[k] = "B"
			}
		}
	}
}

// Get's the unique numbers (1, 4, 7 en 8)
func GetUniques(input string) (uniques []string) {
	numbers := strings.Fields(input)
	var (
		one   string
		four  string
		seven string
		eight string
	)
	for _, pos := range numbers {
		switch len(pos) {
		case 2: // 1
			one = pos
		case 4: // 4
			four = pos
		case 3: // 7
			seven = pos
		case 7: // 8
			eight = pos
		}
	}
	uniques = append(uniques, one, four, seven, eight)
	return uniques
}

// IsIn returns if a is in b
func IsIn(is, in string) bool {
	for _, char := range in {
		if string(char) == is {
			return true
		}
	}
	return false
}

func SilentAtoi(input string) int {
	num, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return num
}
