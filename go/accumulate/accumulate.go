// Package accumulate contains functions for converting string lists
package accumulate

// Accumulate applies the given conversion to the list of strings and returns the modified strings
func Accumulate(collection []string, convert func(string) string) []string {
	for i := range collection {
		collection[i] = convert(collection[i])
	}
	return collection
}
