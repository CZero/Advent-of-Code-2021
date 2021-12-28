package main

import "strconv"

// SilentAtoi returns no errors, only int.
func SilentAtoi(a string) int {
	val, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return val
}
