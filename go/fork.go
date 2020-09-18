package main

import (
	"sync"
)

type Fork struct {
	InUse bool
	mux   sync.Mutex
}

func (fork *Fork) Take() {
	fork.mux.Lock()
	defer fork.mux.Unlock()
	if fork.InUse {
		panic("Taking fork that's in use")
	}
	fork.InUse = true
}

func (fork *Fork) Release() {
	fork.InUse = false
}
