package goroutine

import (
	"errors"
	"testing"
)

func TestRun(t *testing.T) {
	taskList := [...]func() error{
		func() error {
			return nil
		},
		func() error { return errors.New("err") },
		func() error { return errors.New("err") },
		func() error { return errors.New("err") },
		func() error { return errors.New("err") },
		func() error { return errors.New("err") },
		func() error { return errors.New("err") },
		func() error {
			return nil
		},
		func() error {
			return nil
		},
		func() error {
			return nil
		},
		func() error { return errors.New("err") },
		func() error {
			return nil
		},
		func() error { return nil },
		func() error { return nil },
		func() error { return nil },
		func() error { return nil },
		func() error {
			return nil
		},
	}
	tasks := taskList[:]
	err := Run(tasks, 3, 3)

	if err == nil {
		t.Errorf("\n\t%s", "Something goes wrong")
	} else {
		err = Run(tasks, 2, 8)
	}

	if err != nil {
		t.Errorf("\n\t%s", "Something goes wrong")
	} else {
		err = Run(tasks, 4, 7)
	}

	if err == nil {
		t.Errorf("\n\t%s", "Something goes wrong")
	} else {
		err = Run(tasks, 0, 7)
	}

	if err == nil {
		t.Errorf("\n\t%s", "Something goes wrong")
	} else {
		err = Run(tasks, 3, 0)
	}

	if err == nil {
		t.Errorf("\n\t%s", "Something goes wrong")
	}
}
