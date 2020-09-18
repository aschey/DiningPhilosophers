package main

import (
	"fmt"
)

type Waiter struct {
	requestQueue *RequestQueue
}

func NewWaiter() Waiter {
	q := NewRequestQueue()
	return Waiter{requestQueue: q}
}

func (waiter Waiter) Request(philosopher *Philosopher) chan bool {
	requestChan := make(chan bool)
	//defer close(requestChan)
	f := func(name string) {
		fmt.Printf("name: %s closing channel\n", name)
		requestChan <- true
		close(requestChan)
	}
	GetEventMangager().Subscribe(philosopher.Name+"RequestGranted", &f, true)
	waiter.requestQueue.AddRequest(philosopher)
	return requestChan
}
