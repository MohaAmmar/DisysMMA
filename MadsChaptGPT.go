package main

import (
	"fmt"
	"sync"
	"time"
)

const numPhilosophers = 5

type Fork struct{ sync.Mutex }

type Philosopher struct {
	id                       int
	leftFork, rightFork      *Fork
	eatCount                 int
	maxEatingAllowed         int
	eatingTimeMilliseconds   int
	thinkingTimeMilliseconds int
}

func (p *Philosopher) eat() {
	p.leftFork.Lock()
	p.rightFork.Lock()

	fmt.Printf("Philosopher %d is eating for the %d time\n", p.id, p.eatCount+1)

	time.Sleep(time.Duration(p.eatingTimeMilliseconds) * time.Millisecond)

	p.rightFork.Unlock()
	p.leftFork.Unlock()

	p.eatCount++
}

func (p *Philosopher) think() {
	fmt.Printf("Philosopher %d is thinking\n", p.id)
	time.Sleep(time.Duration(p.thinkingTimeMilliseconds) * time.Millisecond)
}

func (p *Philosopher) dine(wg *sync.WaitGroup) {
	defer wg.Done()
	for p.eatCount < p.maxEatingAllowed {
		p.eat()
		p.think()
	}
}

func main() {
	var wg sync.WaitGroup

	forks := make([]*Fork, numPhilosophers)
	philosophers := make([]*Philosopher, numPhilosophers)

	for i := 0; i < numPhilosophers; i++ {
		forks[i] = new(Fork)
	}

	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &Philosopher{
			id:                       i,
			leftFork:                 forks[i],
			rightFork:                forks[(i+1)%numPhilosophers],
			maxEatingAllowed:         3, // Each philosopher will eat 3 times
			eatingTimeMilliseconds:   100,
			thinkingTimeMilliseconds: 100,
		}
		wg.Add(1)
		go philosophers[i].dine(&wg)
	}

	wg.Wait()
}




