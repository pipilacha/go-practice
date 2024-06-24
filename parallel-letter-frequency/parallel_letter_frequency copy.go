package letter

import "sync"

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency2(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// increaseFreq is the go routine that will be increasing the frequency. Only one will be run
// it takes the frequency map, th channel where the other go routines will pass the rune, and the WaitGroup
func increaseFreq2(frequencies FreqMap, ch chan rune, wg *sync.WaitGroup) {
	for r := range ch {
		frequencies[r]++
	}
	wg.Done()
}

// iterateText will read over the text, and write the run it reads to the channel for increaseFreq to increase. Multiple will be run
func iterateText2(text string, ch chan rune, wg *sync.WaitGroup) {
	for _, r := range text {
		ch <- r
	}
	wg.Done()
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency2(texts []string) FreqMap {

	frequencies := FreqMap{}

	freq_channel := make(chan rune, 1000000)

	read_wg := sync.WaitGroup{} // This wait group will be for the routine that increases the frequency. Reads the channel
	read_wg.Add(1)
	go increaseFreq2(frequencies, freq_channel, &read_wg)

	write_wg := sync.WaitGroup{} // This wait group will be for the routines that iterate of the text only. Writes to the channel
	write_wg.Add(len(texts))
	for _, text := range texts {
		go iterateText2(text, freq_channel, &write_wg)
	}
	write_wg.Wait() // Let's wait for the writting routines to finish

	close(freq_channel) // Now that all the writing routines are done let's close the channel so the reading routine finishes

	read_wg.Wait() // Let's wait for the reading routine to finish

	return frequencies
}

/* not good.....
BenchmarkSequentialFrequency-8              5349            227916 ns/op
BenchmarkConcurrentFrequency-8               819           1357931 ns/op
PASS
ok      letter  2.510s
*/
