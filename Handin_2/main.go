package main

import (
	"fmt"
	"math/rand"
	"time"
)

func client(counter int, comms chan Data) {
	var seqC int = rand.Intn(10)
	var ackC int

	if counter == 0 {
		fmt.Printf("seqC is %d \n", seqC)
		data := Data{seqC, 0, "Hello"};
		comms <- data
	} else {
		newData := <- comms
		ackC = newData.seq + 1
		fmt.Printf("ackC is %d \n", ackC)
		fmt.Printf("seqC is %d \n", newData.ack)
	}

}

func server(comms chan Data) {
	var seqS int = rand.Intn(10)
	var ackS int

	data := <-comms 
	ackS = data.seq + 1
	fmt.Printf("ackS is %d \n", ackS)
	fmt.Printf("seqS is %d \n", seqS)
	newData := Data{seqS, ackS, "Hello"};
	comms <- newData
	client(1, comms)
}

type Data struct {
	seq int
	ack int
	msg string
}

func main() {

	var comms = make(chan Data, 2)

	go client(0, comms)
	go server(comms)

	time.Sleep(2 * time.Second)
}
