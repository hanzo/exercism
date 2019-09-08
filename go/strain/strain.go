// Package strain contains functions for filtering elements in a collection
package strain

type Ints []int
type Lists [][]int
type Strings []string

// Keep returns only the values of the given collection that meet the given predicate
func (ints Ints) Keep(predicate func(int) bool) (output Ints) {
	for _, val := range ints {
		if predicate(val) {
			output = append(output, val)
		}
	}
	return
}

// Discard returns only the values of the given collection that do not meet the given predicate
func (ints Ints) Discard(predicate func(int) bool) (output Ints) {
	for _, val := range ints {
		if !predicate(val) {
			output = append(output, val)
		}
	}
	return
}

// Keep returns only the values of the given collection that meet the given predicate
func (lists Lists) Keep(predicate func([]int) bool) (output Lists) {
	for _, val := range lists {
		if predicate(val) {
			output = append(output, val)
		}
	}
	return
}

// Keep returns only the values of the given collection that meet the given predicate
func (strings Strings) Keep(predicate func(string) bool) (output Strings) {
	for _, val := range strings {
		if predicate(val) {
			output = append(output, val)
		}
	}
	return
}
