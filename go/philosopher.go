package main

import "math/rand"

import "time"

import "fmt"

type Philosopher struct {
	Name             string
	LeftFork         *Fork
	RightFork        *Fork
	LeftPhilosopher  *Philosopher
	RightPhilosopher *Philosopher
	Waiter           Waiter
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
	return !philosopher.LeftFork.InUse && !philosopher.RightFork.InUse
}

func (philosopher Philosopher) Think() {
	time.Sleep(time.Duration(philosopher.NextThinkTime()))
}

func (philosopher Philosopher) Eat() {
	requestChan := philosopher.Waiter.Request(philosopher)
	<-requestChan
	philosopher.LeftFork.Take()
	philosopher.RightFork.Take()
	fmt.Printf("%s began eating", philosopher.Name)
	time.Sleep(time.Duration(philosopher.NextEatTime()))
	fmt.Printf("%s finished eating", philosopher.Name)
	philosopher.LeftFork.Release()
	philosopher.RightFork.Release()
}

func (philosopher Philosopher) Run() {
	for true {
		philosopher.Eat()
		philosopher.Think()
	}
}
