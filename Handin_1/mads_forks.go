package main

import (
	"fmt"
)

var numPhilosophers = 5

type Philosopher struct {
	id           int
	mainFork     *Fork
	neighborFork *Fork
}

type Fork struct {
	id          int
	isAvailable chan bool
}

var forks []*Fork

func makePhilosophers(philosopherID int) *Philosopher {
	philosopher := &Philosopher{
		id:           philosopherID,
		mainFork:     forks[(philosopherID-1)%len(forks)], // Wrap around with modulo
		neighborFork: forks[philosopherID%len(forks)],     // Adjust the index
	}

	fmt.Printf("Philosopher %d has mainFork %d and neighborFork %d\n", philosopher.id, philosopher.mainFork.id, philosopher.neighborFork.id)
	return philosopher
}

func makeForks(forkID int) *Fork {
	return &Fork{
		id:          forkID,
		isAvailable: make(chan bool, 1),
	}
}

func initializeForks() {
	for i := 0; i < numPhilosophers; i++ {
		forks[i] = makeForks(i + 1)
	}
}

func philosopher() {
	for i := 0; i < numPhilosophers; i++ {
		makePhilosophers(i + 1)
	}
}

func fork() {
	// Your fork logic here
}

func main() {
	forks = make([]*Fork, numPhilosophers)
	initializeForks()

	go philosopher()
	go fork()

	select {}
}
