// package scale contains functions related to musical scales
package scale

import (
	"fmt"
	"strings"
)

// chromaticScale uses all half steps
const chromaticScale = "mmmmmmmmmmmm"

type Pitch struct {
	sharp_name string
	flat_name  string
}

var pitchList = [12]Pitch{
	Pitch{"A", "A"},
	Pitch{"A#", "Bb"},
	Pitch{"B", "B"},
	Pitch{"C", "C"},
	Pitch{"C#", "Db"},
	Pitch{"D", "D"},
	Pitch{"D#", "Eb"},
	Pitch{"E", "E"},
	Pitch{"F", "F"},
	Pitch{"F#", "Gb"},
	Pitch{"G", "G"},
	Pitch{"G#", "Ab"},
}

// pitchListIndex maps notes (in sharp and flat form) to their index in pitchList
var pitchListIndex = map[string]int{
	"A":  0,
	"A#": 1,
	"Bb": 1,
	"B":  2,
	"C":  3,
	"C#": 4,
	"Db": 4,
	"D":  5,
	"D#": 6,
	"Eb": 6,
	"E":  7,
	"F":  8,
	"F#": 9,
	"Gb": 9,
	"G":  10,
	"G#": 11,
	"Ab": 11,
}

// HalfStepCount returns the number of half steps for the given interval,
// or returns an error if the interval is not recognized
func HalfStepCount(interval rune) (int, error) {
	switch interval {
	case 'm':
		return 1, nil
	case 'M':
		return 2, nil
	case 'A':
		return 3, nil
	default:
		return 0, fmt.Errorf("invalid interval: %v", string(interval))
	}
}

// UseFlatNotes returns true if the given tonic uses flat notes (instead of sharp notes)
func UseFlatNotes(tonic string) bool {
	switch tonic {
	case "F", "Bb", "Eb", "Ab", "Db", "Gb major d", "d", "g", "c", "f", "bb", "eb minor":
		return true
	default:
		return false
	}
}

// PrintPitch returns the string representation of the given pitch, in either sharp or flat form
func PrintPitch(pitch *Pitch, useFlatNotes bool) string {
	if useFlatNotes {
		return pitch.flat_name
	} else {
		return pitch.sharp_name
	}
}

// Scale returns the musical scale starting with the given tonic and following the given interval pattern
func Scale(tonic, intervalPattern string) []string {
	var scale []string
	var curPitch Pitch

	useFlatNotes := UseFlatNotes(tonic)

	curPitchIndex, ok := pitchListIndex[strings.Title(tonic)]
	if !ok {
		fmt.Printf("invalid tonic: %v\n", tonic)
		return []string{}
	}

	// If no interval pattern given, use the chromatic scale
	if intervalPattern == "" {
		intervalPattern = chromaticScale
	}

	for _, interval := range intervalPattern {
		curPitch = pitchList[curPitchIndex]
		scale = append(scale, PrintPitch(&curPitch, useFlatNotes))
		halfSteps, err := HalfStepCount(interval)
		if err != nil {
			fmt.Println(err.Error())
			return []string{}
		}
		// if the current index exceeds the max index, wrap back around to the start of the array
		curPitchIndex = (curPitchIndex + halfSteps) % len(pitchList)
	}
	return scale
}
