package connector

import (
	"github.com/ilaygoldman/go-lib"
)


// Returns the sum of two numbers
func Add2(a int, b int) int {
	return lib.Add(a, b)
}

// Returns the difference between two numbers
func Subtract2(a int, b int) int {
	return lib.Subtract(a, b)
}