package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

const namesMax = 26 * 26 * 10 * 10 * 10 // 676,000
var nameIDs = rand.Perm(namesMax)       // robots will be assigned names from this list, in order
var nextNameIndex = 0                   // index of the nameIDs list to use next

// Robot represents a robot with a unique name
type Robot struct {
	name string
}

// getName converts a given integer ID to the corresponding name. Some examples:
// 0      -> AA000
// 12     -> AA012
// 25000  -> AZ000
// 675999 -> ZZ999
func getName(nameID int) (string, error) {
	if nameID < 0 || nameID >= namesMax {
		return "", errors.New("name ID must be between 0 and 675,999 (inclusive)")
	}
	idStr := fmt.Sprintf("%06d", nameID) // left-pad with zeroes to 6 digits (since max is 675,999)

	// the first three digits of the integer correspond to the first two letters of the name.
	// convert these from base 10 to base 26
	letterIntVal, err := strconv.Atoi(idStr[0:3])
	if err != nil {
		return "", err
	}
	letter1 := letterIntVal / 26
	letter2 := letterIntVal % 26

	var builder strings.Builder

	builder.WriteRune('A' + rune(letter1))
	builder.WriteRune('A' + rune(letter2))
	builder.WriteByte(idStr[3])
	builder.WriteByte(idStr[4])
	builder.WriteByte(idStr[5])

	return builder.String(), nil
}

// getNextName returns the next unused randomly generated name in the sequence.
func getNextName() (string, error) {
	if nextNameIndex >= namesMax {
		return "", errors.New("we're all out of names")
	}
	name, err := getName(nameIDs[nextNameIndex])
	if err != nil {
		return "", err
	}
	nextNameIndex++
	return name, nil
}

// Reset sets the robot's name to a different (and unused) name.
func (r *Robot) Reset() error {
	name, err := getNextName()
	if err != nil {
		return err
	}
	r.name = name
	return nil
}

// Name returns the robot's name if it has already been set, otherwise
// the robot's name is set to a random value and this value is returned.
func (r *Robot) Name() (string, error) {
	if len(r.name) > 0 {
		return r.name, nil
	}
	name, err := getNextName()
	if err != nil {
		return "", err
	}
	r.name = name
	return name, nil
}
