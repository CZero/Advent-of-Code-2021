package main

import (
	"fmt"
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

 **/

func main() {
	// lines, _ := readLines("example.txt")
	lines, _ := readLines("input.txt")
	fmt.Println(FindUniqueDisplayNumbers(lines))
}

func FindUniqueDisplayNumbers(input []string) (found int) {
	for _, line := range input {
		parts := strings.Split(line, "|") // 0 is digits, 1 is display
		display := strings.Fields(parts[1])
		for _, pos := range display {
			switch len(pos) {
			case 2: // 1
				found++
			case 4: // 4
				found++
			case 3: // 7
				found++
			case 7: // 8
				found++
			}

		}
	}
	return found
}
