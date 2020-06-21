package letter

// FreqMap records the frequency of each
// rune in a given text.
type FreqMap map[rune]int

var freqChan = make(chan FreqMap)

// Frequency counts the frequency of each
// rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency does the same thing
// as Frequency, but using Go routines
func ConcurrentFrequency(strings []string) FreqMap {
	m := FreqMap{}

	for _, s := range strings {
		go func(str string) {
			freqChan <- Frequency(str)
		}(s)
	}

	for range strings {
		for key, value := range <-freqChan {
			m[key] += value
		}
	}

	return m
}
