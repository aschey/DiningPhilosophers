package main

import "fmt"

type RequestQueue struct {
	requests        *PriorityQueue
	requestNames    *ConcurrentHashSet
	pendingRequests *ConcurrentHashSet
	maxRequestNames int
}

type Request struct {
	MaxPriority int
	Philosopher *Philosopher
	Priority    int
}

func (request Request) Overdue() bool {
	return request.Priority >= request.MaxPriority
}

func NewRequestQueue() RequestQueue {
	r := NewConcurrentHashSet()
	p := NewConcurrentHashSet()
	requestQueue := RequestQueue{
		requestNames:    &r,
		pendingRequests: &p,
		maxRequestNames: 10,
		requests:        new(PriorityQueue),
	}

	eventMangager := GetEventMangager()
	f1 := func(name string) {
		requestQueue.pendingRequests.Remove(name)
	}
	eventMangager.Subscribe("Finished",
		&f1,
		false)
	f2 := func(_ string) {
		requestQueue.Run()
	}
	eventMangager.Subscribe("RequestAdded",
		&f2,
		true)

	return requestQueue
}

func (requestQueue RequestQueue) Count() int {
	return requestQueue.requests.Len()
}

func (requestQueue *RequestQueue) AddRequest(philosopher *Philosopher) {
	requestQueue.requestNames.Add(philosopher.Name)
	requestQueue.requests.Push(&Item{value: Request{Philosopher: philosopher}, priority: 0})
	GetEventMangager().Broadcast("RequestAdded", "")
}

func (requestQueue RequestQueue) Run() {
	for requestQueue.Count() > 0 {
		requestVal := requestQueue.requests.Pop()
		requestItem := requestVal.(*Item)
		request := requestItem.value.(Request)
		philosopher := request.Philosopher

		leftNeighborRequested := requestQueue.requestNames.Contains(philosopher.LeftPhilosopher.Name)
		rightNeighborRequested := requestQueue.requestNames.Contains(philosopher.RightPhilosopher.Name)
		lessThanTwoNeighborsRequested := !(leftNeighborRequested && rightNeighborRequested)
		leftNeighorGranted := requestQueue.pendingRequests.Contains(philosopher.LeftPhilosopher.Name)
		rightNeighborGranted := requestQueue.pendingRequests.Contains(philosopher.RightPhilosopher.Name)
		neighborGranted := leftNeighorGranted || rightNeighborGranted

		if philosopher.CanEat() && !neighborGranted && (request.Overdue() || lessThanTwoNeighborsRequested || requestQueue.requestNames.Length() > requestQueue.maxRequestNames) {
			requestQueue.requestNames.Remove(philosopher.Name)
			requestQueue.pendingRequests.Add(philosopher.Name)
			fmt.Printf("granting %s", philosopher.Name)
			GetEventMangager().Broadcast(philosopher.Name+"RequestGranted", "")
		} else {
			requestItem.priority++
			requestQueue.requests.Push(requestItem)

		}
	}
	f := func(_ string) { requestQueue.Run() }
	GetEventMangager().Subscribe("RequestAdded", &f, true)
}
