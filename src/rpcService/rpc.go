package rpcService

import (
	"errors"
	"time"
)

type Args struct {
	A, B int
}

type Calculate int

func (t *Calculate)Add(a1 Args, result *int) error {

	*result = a1.A + a1.B
	time.Sleep(5*time.Second)

	return nil
}

func (t *Calculate)Devide(a1 Args, result *float64) error {
	if a1.B == 0 {
		return errors.New("divide by zero")
	}
	*result = float64(a1.A / a1.B)
	return nil
}


