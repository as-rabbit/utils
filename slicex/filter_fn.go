package slicex

import (
	"golang.org/x/exp/constraints"
)

// StringEmpty  == ""
func StringEmpty[T string](i int, item T) bool {
	if len(item) > 0 {
		return true
	}
	return false
}

// IntegerEmpty  == 0
func IntegerEmpty[T constraints.Integer | constraints.Float](i int, item T) bool {
	if item != 0 {
		return true
	}
	return false
}
