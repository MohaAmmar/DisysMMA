package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Fork struct {
	id       int
	forkChan chan bool
}

func newFork(id int, forkChan chan bool) *Fork {
	return &Fork{id, forkChan}
}

func (fork *Fork) listen() {
	for true {
		select {
		case <-fork.forkChan:
			fork.forkChan <- true
		}
		time.Sleep(time.Duration(rand.Int63n(1e9)))
	}
}

type Philosopher struct {
	name                string
	leftFork, rightFork chan bool
	eatCounter          int
}

func newPhilosopher(name string, leftForkChan, rightForkChan chan bool) *Philosopher {
	return &Philosopher{name, leftForkChan, rightForkChan, 0}
}

func (phil *Philosopher) dine(wg *sync.WaitGroup) {
	fmt.Printf("%v is now dining \n", phil.name)

	defer wg.Done()
	for phil.eatCounter < 3 {
		phil.think()
		phil.getForks()
		phil.eat()
		phil.returnForks()
	}
	fmt.Printf("%v has finished dining! \n", phil.name)
}

func (phil *Philosopher) getForks() {
	timeout := make(chan bool, 1)
	go func() { time.Sleep(1e9); timeout <- true }()

	<-phil.leftFork
	select {
	case <-phil.rightFork:
		return
	case <-timeout:
		phil.leftFork <- true
		phil.think()
		phil.getForks()
	}
}

func (phil *Philosopher) returnForks() {
	phil.leftFork <- true
	phil.rightFork <- true
}

func (phil *Philosopher) eat() {
	phil.eatCounter++
	fmt.Printf("%v is eating for the %d time! \n", phil.name, phil.eatCounter)
	time.Sleep(time.Duration(rand.Int63n(1e9)))
}

func (phil *Philosopher) think() {
	fmt.Printf("%v is thinking \n", phil.name)
	time.Sleep(time.Duration(rand.Int63n(1e9)))
}

func main() {

	const numPhilosophers = 5
	var wg sync.WaitGroup
	names := []string{"Phil 1", "Phil 2", "Phil 3", "Phil 4", "Phil 5"}

	var forkChannels [numPhilosophers]chan bool
	for i := 0; i < numPhilosophers; i++ {
		forkChannels[i] = make(chan bool, 1)
		forkChannels[i] <- true
	}

	var forks [numPhilosophers]*Fork
	var philosophers [numPhilosophers]*Philosopher
	for i, name := range names {
		forks[i] = newFork(i, forkChannels[i])
		go forks[i].listen()
		philosophers[i] = newPhilosopher(name, forkChannels[(i)], forkChannels[((i+1)%5)])
	}

	for i := 0; i < numPhilosophers; i++ {
		wg.Add(1)
		go philosophers[i].dine(&wg)
	}

	wg.Wait()
}
