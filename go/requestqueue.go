package main

type RequestQueue struct {
	requests        PriorityQueue
	requestNames    *ConcurrentHashSet
	pendingRequests *ConcurrentHashSet
	maxRequestNames int
}

type Request struct {
	MaxPriority int
	Philosopher Philosopher
	Priority    int
}

func (request Request) Overdue() bool {
	return request.Priority >= request.MaxPriority
}

func NewRequestQueue() *RequestQueue {
	requestQueue := new(RequestQueue)
	requestQueue.requestNames = NewConcurrentHashSet()
	requestQueue.pendingRequests = NewConcurrentHashSet()
	requestQueue.maxRequestNames = 10

	eventMangager := GetEventMangager()
	eventMangager.Subscribe("Finished",
		func(name string) { requestQueue.pendingRequests.Remove(name) },
		false)

	eventMangager.Subscribe("RequestAdded",
		func(_ string) { requestQueue.Run() },
		true)

	return requestQueue
}

func (requestQueue RequestQueue) AddRequest(philosopher Philosopher) {
	requestQueue.requestNames.Add(philosopher.Name)
	request := new(Request)
	request.Philosopher = philosopher
	//requestQueue.requests.Enqueue()
}

func (requestQueue RequestQueue) Run() {

}
