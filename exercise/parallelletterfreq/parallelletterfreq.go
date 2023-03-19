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

// ConcurrentFrequency counts the frequency of each rune in a given text concurrently
func ConcurrentFrequency(strSlice []string) FreqMap {
	cm := make(chan FreqMap, len(strSlice))
	cs := make(chan string, len(strSlice))
	var wg sync.WaitGroup
	for i, s := range strSlice {
		wg.Add(1)
		go func(cm chan<- FreqMap, cs <-chan string) {
			innerS := <-cs
			cm <- Frequency(innerS)
			wg.Done()
		}(cm, cs)
		cs <- s
		if i == len(strSlice)-1 {
			close(cs)
		}
	}
	go func() {
		wg.Wait()
		close(cm)
	}()
	result := FreqMap{}
	for d := range cm {
		for k, v := range d {
			result[k] += v
		}
	}
	return result
}
