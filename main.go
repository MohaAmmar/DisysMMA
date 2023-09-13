package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const numPhilosophers = 5

type Fork struct {
	sync.Mutex
	id int
}

type Philosopher struct {
	id                  int
	leftFork, rightFork chan bool
}

func (p *Philosopher) eat() {
	n := rand.Intn(10)
	time.Sleep(time.Duration(n) * time.Second)
	<-p.leftFork
	<-p.rightFork

	fmt.Printf("Philosopher %d is eating\n", p.id)

	p.leftFork <- true
	p.rightFork <- true
}

func (p *Philosopher) think() {
	n := rand.Intn(10)
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Printf("Philosopher %d is thinking\n", p.id)
}

func forkRoutine(fork *Fork, eatCh chan bool) {
	for {
		eatCh <- true // Wait for signal to pick up fork
		fork.Lock()
		fmt.Printf("Fork %d is picked up\n", fork.id)
		<-eatCh // Wait for signal to put down fork
		fork.Unlock()
		fmt.Printf("Fork %d is put down\n", fork.id)
	}
}

func philosopherRoutine(id int, leftFork, rightFork chan bool, doneCh chan bool) {
	philosopher := Philosopher{id, leftFork, rightFork}
	for i := 0; i < 3; i++ {
		philosopher.think()
		philosopher.eat()
		fmt.Printf("%d has eaten %d times\n", philosopher.id, i)
	}
	doneCh <- true
}

func main() {
	forks := make([]chan bool, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		forks[i] = make(chan bool)
		fork := &Fork{id: i}
		go forkRoutine(fork, forks[i])
	}

	doneCh := make(chan bool, numPhilosophers)

	for i := 0; i < numPhilosophers; i++ {
		go philosopherRoutine(i, forks[i], forks[(i+1)%numPhilosophers], doneCh)
	}

	for i := 0; i < numPhilosophers; i++ {
		<-doneCh
	}
}
