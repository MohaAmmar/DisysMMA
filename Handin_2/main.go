package main

import (
	"fmt"
	"math/rand"
	"time"
)

func client(comms chan int) {
	var counter int = 0
	var seqC int = rand.Intn(10)
	var ackC int

	/*if counter == 0 {
		clientFirstListen(seqC, comms, wg)
		counter++
	} else {
		clientListen(seqC, ackC, comms, wg)
	}*/

	if counter == 0 {
		fmt.Printf("seqC is %d \n", seqC)
		comms <- seqC
		counter++
	} else {
		<-comms
		ackC = <-comms
		seqC = <-comms
		fmt.Printf("ackC is %d \n", ackC)
		fmt.Printf("seqC is %d \n", seqC)
		<-comms
		comms <- ackC
		comms <- seqC
	}

}

func server(comms chan int) {
	var seqS int = rand.Intn(10)
	var ackS int

	ackS = <-comms
	fmt.Printf("ackS is %d \n", ackS)
	fmt.Printf("seqS is %d \n", seqS)
	comms <- ackS
	comms <- seqS
}

/*func clientFirstListen(seqC int, comms chan int, wg *sync.WaitGroup) {
	comms <- seqC
}

func clientListen(seqC int, ackC int, comms chan int, wg *sync.WaitGroup) (int, int){
		ackC := <-comms
		seqC := <-comms
		<-comms
		comms <- ackC
		comms <- seqC
	return ackC, seqC
}

func serverListen(seqS int, ackS int, comms chan int, wg *sync.WaitGroup) {

}*/

func main() {

	var comms = make(chan int, 2)

	go client(comms)
	go server(comms)

	time.Sleep(100 * time.Second)
}
