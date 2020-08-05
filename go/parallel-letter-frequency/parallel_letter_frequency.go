package letter

import (
	"sync"
)

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

type SafeFreqMap struct {
	Map map[rune]int
	mux sync.Mutex
}

func New() *SafeFreqMap {
	return &SafeFreqMap{
		Map: make(map[rune]int),
	}
}

func (s *SafeFreqMap) Increment(r rune) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.Map[r]++
}

func SafeFrequency(s string, m *SafeFreqMap) {
	for _, r := range s {
		m.Increment(r)
	}
}

func ConcurrentFrequency(strs []string) FreqMap {
	// // works, but serial:
	// m := FreqMap{}
	// for _, s := range strs {
	// 	subMap := Frequency(s)
	// 	for key, val := range subMap {
	// 		m[key] += val
	// 	}
	// }
	// sfm := SafeFreqMap{}

	// // works, but dumb:
	// sfm := New()
	// for _, s := range strs {
	// 	go SafeFrequency(s, sfm)
	// }
	// time.Sleep(1 * time.Second)
	// m := FreqMap{}
	// for k, v := range sfm.Map {
	// 	m[k] = v
	// }

	// // works using wait group and custom struct
	// sfm := New()
	// var wg sync.WaitGroup
	// wg.Add(len(strs))
	// for _, s := range strs {
	// 	go func(s string) {
	// 		SafeFrequency(s, sfm)
	// 		wg.Done()
	// 	}(s)
	// }
	// wg.Wait()
	// m := FreqMap{}
	// for k, v := range sfm.Map {
	// 	m[k] = v
	// }

	// // works using wait group, but requires iterating over each FreqMap
	// freqMaps := make([]FreqMap, len(strs))
	// var wg sync.WaitGroup
	// wg.Add(len(strs))
	// for i, s := range strs {
	// 	go func(i int, s string) {
	// 		freqMaps[i] = Frequency(s)
	// 		wg.Done()
	// 	}(i, s)
	// }
	// wg.Wait()
	// m := FreqMap{}
	// for _, subMap := range freqMaps {
	// 	for k, v := range subMap {
	// 		m[k] += v
	// 	}
	// }
	// return m

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
