package main

type Waiter struct {
	requestQueue RequestQueue
}

func (waiter Waiter) Request(philosopher Philosopher) chan bool {
	requestChan := make(chan bool)
	defer close(requestChan)
	GetEventMangager().Subscribe(philosopher.Name+"RequestGranted", func(name string) {
		requestChan <- true
	}, true)
	waiter.requestQueue.AddRequest(philosopher)
	return requestChan
}
