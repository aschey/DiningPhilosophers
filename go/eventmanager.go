package main

import "sync"

import cmap "github.com/orcaman/concurrent-map"

type EventArgs struct {
	Callbacks       []func(string)
	AutoUnsubscribe bool
	Mux             sync.Mutex
}

type EventManager struct {
	eventHandlers cmap.ConcurrentMap
}

func (manager EventManager) Broadcast(eventName string, arg string) {
	value, _ := manager.eventHandlers.Get(eventName)
	eventArgs, _ := value.(EventArgs)
	eventArgs.Mux.Lock()
	for _, callback := range eventArgs.Callbacks {
		go callback(arg)
	}
	if eventArgs.AutoUnsubscribe {
		eventArgs.Callbacks = nil
	}
	eventArgs.Mux.Unlock()
}

func (manager EventManager) Subscribe(name string, callback func(string), autoUnsubscribe bool) {
	value, contains := manager.eventHandlers.Get(name)
	if contains {
		args, _ := value.(EventArgs)
		args.Mux.Lock()
		args.AutoUnsubscribe = autoUnsubscribe
		args.Callbacks = append(args.Callbacks, callback)
		args.Mux.Unlock()
	} else {
		args := EventArgs()
		args.AutoUnsubscribe = autoUnsubscribe
		args.Callbacks = [1]func(string){callback}
		manager.eventHandlers.Set(name, args)
	}
}
