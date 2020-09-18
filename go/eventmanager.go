package main

import (
	"sync"

	cmap "github.com/orcaman/concurrent-map"
)

type EventArgs struct {
	Callbacks       []*func(string)
	AutoUnsubscribe bool
	Mux             sync.Mutex
}

type EventManager struct {
	eventHandlers cmap.ConcurrentMap
}

var (
	eventManagerSingleton *EventManager
	once                  sync.Once
)

func GetEventMangager() *EventManager {
	once.Do(func() {
		eventManagerSingleton = &EventManager{eventHandlers: cmap.New()}
	})

	return eventManagerSingleton
}

func (manager *EventManager) Broadcast(eventName string, arg string) {
	value, _ := manager.eventHandlers.Get(eventName)
	eventArgs, _ := value.(*EventArgs)
	eventArgs.Mux.Lock()
	defer eventArgs.Mux.Unlock()
	for _, callback := range eventArgs.Callbacks {
		cFunc := *callback
		go cFunc(arg)
	}
	if eventArgs.AutoUnsubscribe {
		eventArgs.Callbacks = nil
	}

}

func (manager *EventManager) Subscribe(name string, callback *func(string), autoUnsubscribe bool) {
	value, contains := manager.eventHandlers.Get(name)
	if contains {
		args, _ := value.(*EventArgs)
		args.Mux.Lock()
		defer args.Mux.Unlock()
		args.AutoUnsubscribe = autoUnsubscribe
		args.Callbacks = append(args.Callbacks, callback)

	} else {
		args := new(EventArgs)
		args.AutoUnsubscribe = autoUnsubscribe
		args.Callbacks = []*func(string){callback}
		manager.eventHandlers.Set(name, args)
	}
}
