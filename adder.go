// Solution for excercises in chaper 10 of book Learning Go, 2nd edition.
package Adder

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Add is a function that takes two integers and adds them together, returning their sum.
func Add[T Number](n1, n2 T) T {
	return n1 + n2
}
