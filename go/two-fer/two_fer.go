// Package twofer contains functions for sharing.
package twofer

import "fmt"

// ShareWith returns a two-fer string for the given name.
// If no name is provided, 'you' is used as the name.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
