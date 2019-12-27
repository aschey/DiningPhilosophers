package main

import "sync"

type Table struct {
	Philosophers []*Philosopher
	Forks        []*Fork
	Waiter       *Waiter
}

func NewTable() Table {
	names := []string{
		"Aristotle",
		"Socrates",
		"Confucius",
		"Newton",
		"Locke",
		"Kant",
		"Marx",
		"Nietzsche",
		"Darwin",
		"Descartes",
		"Machiavelli",
		"Hobbes",
		"Chomsky",
	}

	forks := make([]*Fork, len(names))
	for i := range forks {
		forks[i] = &Fork{}
	}
	philosophers := make([]*Philosopher, len(names))
	waiter := NewWaiter()

	for i := range names {
		nextFork := i + 1
		if nextFork == len(names) {
			nextFork = 0
		}
		philosopher := Philosopher{
			Name:          names[i],
			LeftFork:      forks[i],
			RightFork:     forks[nextFork],
			Waiter:        &waiter,
			ThinkTime:     0,
			EatTime:       0,
			ThinkVariance: 0,
			EatVariance:   0,
		}
		philosophers[i] = &philosopher
	}

	for i := range names {
		left := i - 1
		right := i + 1
		if left == -1 {
			left = len(names) - 1
		}
		if right == len(names) {
			right = 0
		}
		philosophers[i].LeftPhilosopher = philosophers[left]
		philosophers[i].RightPhilosopher = philosophers[right]
	}

	table := Table{
		Forks:        forks,
		Philosophers: philosophers,
		Waiter:       &waiter,
	}

	return table
}

func (table Table) Run() {
	wg := sync.WaitGroup{}
	for _, philosopher := range table.Philosophers {
		wg.Add(1)
		go philosopher.Run(&wg)
	}
	wg.Wait()
}
