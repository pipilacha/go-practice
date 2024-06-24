package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {

	frequencies := FreqMap{}

	freq_channel := make(chan FreqMap) // channel where we will store the frequencies returned by each routine

	for _, text := range texts {
		go func(text string) {
			freq_channel <- Frequency(text) // we will pass the result of Frequency to the channel
		}(text)
	}

	/*
		The routinesd will lock until the channel is read and they can write again.
		Since we know the amount of routines = len(texts) we know that every routine writes one time
		then we now how many times we have to read the channel
	*/

	for range texts { // read the channel len(texts) times
		for k, v := range <-freq_channel {
			frequencies[k] += v
		}

	}

	return frequencies
}

/*
now we have faster times

go test -bench .
BenchmarkSequentialFrequency-8              5228            225756 ns/op
BenchmarkConcurrentFrequency-8             10129            116891 ns/op
PASS
ok      letter  3.266s

go test -benchmem -run=^$ -bench ^Benchmark* letter
BenchmarkSequentialFrequency-8              5193            230770 ns/op           17567 B/op         13 allocs/op
BenchmarkConcurrentFrequency-8             10327            116830 ns/op           12588 B/op         70 allocs/op
PASS
ok      letter  3.208s
*/
