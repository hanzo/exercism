package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency concurrently counts the frequency of each rune
// in the given list of texts and returns the merged data as a FreqMap.
func ConcurrentFrequency(strs []string) FreqMap {
	result := FreqMap{}
	ch := make(chan FreqMap)
	for _, s := range strs {
		go func(s string) {
			ch <- Frequency(s)
		}(s)
	}

	for range strs {
		submap := <-ch
		for k, v := range submap {
			result[k] += v
		}
	}
	return result
}
