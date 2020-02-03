package goroutine

import (
	"errors"
	"testing"
)

func TestRun(t *testing.T) {
	testInt := 0
	taskList := [...]func() error{
		func() error {
			testInt = 1
			return nil
		},
		func() error { return errors.New("err") },
		func() error { return errors.New("err") },
		func() error { return errors.New("err") },
		func() error { return errors.New("err") },
		func() error { return errors.New("err") },
		func() error { return errors.New("err") },
		func() error {
			testInt = 1
			return nil
		},
		func() error {
			testInt = 1
			return nil
		},
		func() error {
			testInt = 2
			return nil
		},
		func() error { return errors.New("err") },
		func() error {
			testInt = 2
			return nil
		},
		func() error { return nil },
		func() error { return nil },
		func() error { return nil },
		func() error { return nil },
		func() error {
			testInt = 3
			return nil
		},
	}
	tasks := taskList[:]
	Run(tasks, 3, 6)

	if testInt != 1 {
		t.Errorf("\n\t%s", "Something goes wrong")
	} else {
		Run(tasks, 2, 8)
	}

	if testInt != 3 {
		t.Errorf("\n\t%s", "Something goes wrong")
	} else {
		Run(tasks, 4, 7)
	}

	if testInt != 2 {
		t.Errorf("\n\t%s", "Something goes wrong")
	}
}
