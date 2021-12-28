// uls: Usefull Library Steph
package uls

import "strconv"

// SilentAtoi returns an int from a string and panics when it fails.
// Added during AoC2021
func SilentAtoi(input string) int {
	val, err := strconv.Atoi(input)
	if err != nil {
		panic("SilentAtoi: This is not a number")
	}
	return val
}
