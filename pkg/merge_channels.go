package pkg

func MergeChannels[T any](outChan chan T, inChans ...<-chan T) {
	for _, inChan := range inChans {
		go func(c <-chan T) {
			for v := range c {
				outChan <- v
			}
		}(inChan)
	}
	return
}
