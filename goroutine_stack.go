package goroutine

import (
	"errors"
	"sync"
)

//Run to start execution of functions in tasks slice in N goroutines until maxErrCount errors
func Run(tasks []func() error, N int, maxErrCount int) error {
	if N <= 0 || maxErrCount <= 0 {
		return errors.New("Wrong parameters")
	}

	var goroutCount int
	if N < len(tasks) {
		goroutCount = N
	} else {
		goroutCount = len(tasks)
	}

	mur := sync.RWMutex{}
	errCount := 0
	var (
		wg       sync.WaitGroup
		funcChan = make(chan func() error, len(tasks))
	)
	for i := 1; i <= goroutCount; i++ {
		wg.Add(1)
		go func() {
			for currentFunc := range funcChan {
				if checkErrCount(&errCount, maxErrCount, &mur) {
					break
				}
				err := currentFunc()
				if err != nil {
					incErrCount(&errCount, &mur)
				}
			}
			wg.Done()
		}()
	}
	for _, task := range tasks {
		funcChan <- task
	}
	close(funcChan)
	wg.Wait()
	if checkErrCount(&errCount, maxErrCount, &mur) {
		return errors.New("Too much errors")
	}
	return nil
}

func checkErrCount(errCount *int, compare int, mur *sync.RWMutex) bool {
	mur.RLock()
	defer mur.RUnlock()
	if *errCount < compare {
		return false
	}
	return true
}

func incErrCount(errCount *int, mur *sync.RWMutex) {
	mur.Lock()
	defer mur.Unlock()
	*errCount++
}
