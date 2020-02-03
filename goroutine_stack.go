package goroutine

//Run to start execution of functions in tasks slice in N goroutines until M errors
func Run(tasks []func() error, N int, M int) error {
	var goroutCount int
	if N < len(tasks) {
		goroutCount = N
	} else {
		goroutCount = len(tasks)
	}

	var ErrorChan = make(chan error, M)
	var FuncChan = make(chan func() error)
	var CheckChan = make(chan bool, goroutCount)
	var StopChan = make(chan bool)
	for i := 1; i <= goroutCount; i++ {
		go func(i int) {
			for currentFunc := range FuncChan {
				if !checkErr(ErrorChan) {
					err := currentFunc()
					if err != nil {
						ErrorChan <- err
						if checkErr(ErrorChan) {
							StopChan <- true
						}
					}
				}
			}
			if !checkErr(ErrorChan) {
				CheckChan <- true
				if len(CheckChan) == goroutCount {
					StopChan <- true
				}
			}
		}(i)
	}
	for _, task := range tasks {
		FuncChan <- task
	}
	close(FuncChan)
	<-StopChan
	close(StopChan)
	close(CheckChan)
	close(ErrorChan)
	return nil
}

func checkErr(ErrorChan chan error) bool {
	if len(ErrorChan) < cap(ErrorChan) {
		return false
	} else {
		return true
	}
}
