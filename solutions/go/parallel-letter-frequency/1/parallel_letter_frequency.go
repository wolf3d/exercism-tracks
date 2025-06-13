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

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	var wg sync.WaitGroup
	var c chan FreqMap = make(chan FreqMap)
	var r FreqMap = make(FreqMap)

	wg.Add(len(l))

	go func() {
		wg.Wait()
		close(c)
	}()

	g := func(wg *sync.WaitGroup, word string) {
		defer wg.Done()
		c <- Frequency(word)
	}

	for _, word := range l {
		go g(&wg, word)
	}

	for dictionary := range c {
		for i := range dictionary {
			r[i] += dictionary[i]
		}
	}

	return r
}
