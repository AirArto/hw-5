package goroutine

import (
	"sync"
)

//Run to start execution of functions in tasks slice in N goroutines until M errors
func Run(tasks []func() error, N int, M int) error {
	var goroutCount int
	if N < len(tasks) && N > 0 {
		goroutCount = N
	} else {
		goroutCount = len(tasks)
	}

	var errCount int
	if M > 0 {
		errCount = M
	} else {
		M = 0
		errCount = goroutCount
	}

	mur := sync.RWMutex{}
	var once sync.Once
	var errorChan = make(chan error, errCount)
	var funcChan = make(chan func() error, goroutCount)
	var checkChan = make(chan bool, goroutCount)
	var stopChan = make(chan bool)
	for i := 1; i <= goroutCount; i++ {
		go func() {
			for currentFunc := range funcChan {
				if compareChanLen(errorChan, &mur, errCount) {
					continue
				}
				err := currentFunc()
				if err != nil {
					if !compareChanLen(errorChan, &mur, errCount) {
						mur.Lock()
						errorChan <- err
						mur.Unlock()
					}
					if compareChanLen(errorChan, &mur, errCount) {
						once.Do(func() { stopChan <- true })
					}
				}
			}
			if !compareChanLen(errorChan, &mur, errCount) {
				checkChan <- true
				if len(checkChan) == goroutCount {
					once.Do(func() { stopChan <- true })
				}
			}
		}()
	}
	for _, task := range tasks {
		funcChan <- task
	}
	close(funcChan)
	<-stopChan
	close(stopChan)
	close(checkChan)
	close(errorChan)
	return nil
}

// Compared len of chan with second parameter
func compareChanLen(someChan chan error, mur *sync.RWMutex, compare int) bool {
	mur.RLock()
	defer mur.RUnlock()
	if compare == 0 {
		return false
	}
	if len(someChan) < compare {
		return false
	}
	return true
}
