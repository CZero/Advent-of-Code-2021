// lfs stands for Lib Functions Steph
package lfs

import "strconv"

// SilentAtoi returns an int or panics
func SilentAtoi(input string) int {
	val, err := strconv.Atoi(input)
	PanErr(err)
	return val
}
