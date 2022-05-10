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

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	// This needs to be a buffered channel because otherwise writers will wait until a value has
	// been read before writing, thus losing some benefits of concurrency.
	channel := make(chan FreqMap, len(l))
	output := FreqMap{}
	for _, s := range l {
		go countLettersInString(channel, s)
		if len(output) > 0 {
			m := <-channel
			for k, v := range m {
				output[k] += v
			}
		} else {
			output = <-channel
		}
	}

	return output
}

func countLettersInString(c chan<- FreqMap, s string) {
	c <- Frequency(s)
}
