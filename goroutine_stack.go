package goroutine

import "fmt"

//Run to start execution of functions in tasks slice in N goroutines until M errors
func Run(tasks []func() error, N int, M int) error {
	var goroutCount int
	var bufferSize int
	if N < len(tasks) {
		goroutCount = N
		bufferSize = len(tasks)
	} else {
		goroutCount = len(tasks)
		bufferSize = N
	}

	var ErrorChan = make(chan error, M)
	var FuncChan = make(chan func() error, bufferSize)
	var CheckChan = make(chan bool, goroutCount)
	var StopChan = make(chan bool)
	for i := 0; i < goroutCount; i++ {
		go func(i int) {
			fmt.Printf("goroutine %v started", i)
			for currentFunc := range FuncChan {
				if len(ErrorChan) == cap(ErrorChan) {
					close(FuncChan)
					break
				}
				err := currentFunc()
				if err != nil {
					ErrorChan <- err
				}

			}
			CheckChan <- true
			if len(CheckChan) == goroutCount {
				StopChan <- true
			}
			fmt.Printf("goroutine %v stoped", i)
		}(i)
	}
	for _, task := range tasks {
		FuncChan <- task
	}
	<-StopChan
	fmt.Println("all goro stoped")
	return nil
}
