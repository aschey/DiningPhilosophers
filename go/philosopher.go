package main

import "math/rand"

import "time"

import "fmt"

import "sync"

type Philosopher struct {
	Name             string
	LeftFork         *Fork
	RightFork        *Fork
	LeftPhilosopher  *Philosopher
	RightPhilosopher *Philosopher
	Waiter           *Waiter
	ThinkTime        int
	EatTime          int
	ThinkVariance    int
	EatVariance      int
}

func getVariance(baseTime int, variance int) int {
	return baseTime + int(float32(-variance)+rand.Float32()*2.0*float32(variance))
}

func (philosopher Philosopher) NextThinkTime() int {
	return getVariance(philosopher.ThinkTime, philosopher.ThinkVariance)
}

func (philosopher Philosopher) NextEatTime() int {
	return getVariance(philosopher.EatTime, philosopher.EatVariance)
}

func (philosopher Philosopher) CanEat() bool {
	fmt.Printf("%s can eat: %t\n", philosopher.Name, !philosopher.LeftFork.InUse && !philosopher.RightFork.InUse)
	return !philosopher.LeftFork.InUse && !philosopher.RightFork.InUse
}

func (philosopher Philosopher) Think() {
	fmt.Printf("%s thinking\n", philosopher.Name)
	time.Sleep(time.Duration(philosopher.NextThinkTime()))
}

func (philosopher *Philosopher) Eat() {
	fmt.Printf("%s awaiting eat\n", philosopher.Name)
	requestChan := philosopher.Waiter.Request(philosopher)
	res := <-requestChan
	fmt.Printf("granted %t\n", res)
	philosopher.LeftFork.Take()
	philosopher.RightFork.Take()
	fmt.Printf("%s began eating\n", philosopher.Name)
	time.Sleep(time.Duration(philosopher.NextEatTime()))
	fmt.Printf("%s finished eating\n", philosopher.Name)
	philosopher.LeftFork.Release()
	philosopher.RightFork.Release()
	GetEventMangager().Broadcast("Finished", philosopher.Name)
}

func (philosopher *Philosopher) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	for true {
		philosopher.Eat()
		philosopher.Think()
	}
}
