package main

import (
	"errors"
	"sync"
)

type Fork struct {
	InUse bool
	mux   sync.Mutex
}

func (fork *Fork) Take() error {
	fork.mux.Lock()
	if fork.InUse {
		return errors.New("Taking fork that's in use")
	}
	fork.InUse = true
	return nil
}

func (fork *Fork) Release() {
	fork.InUse = false
}
