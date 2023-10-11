package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Data struct {
	seq int
	ack int
	msg string
}

func client(comms chan Data) {
	var seqC int = rand.Intn(10)
	var ackC int

	fmt.Printf("seqC is %d \n\n", seqC)
	firstShake := Data{seqC, 0, ""}
	comms <- firstShake

	time.Sleep(time.Second / 2)

	secondShake := <-comms
	ackC = secondShake.seq + 1
	fmt.Printf("ackC is %d \n", ackC)
	fmt.Printf("seqC is %d \n\n", secondShake.ack)
	thirdShake := Data{secondShake.ack, ackC, "Hello"}
	comms <- thirdShake

}

func server(comms chan Data) {
	var seqS int = rand.Intn(10)
	var ackS int

	firstShake := <-comms
	ackS = firstShake.seq + 1
	fmt.Printf("ackS is %d \n", ackS)
	fmt.Printf("seqS is %d \n\n", seqS)
	secondShake := Data{seqS, ackS, ""}
	comms <- secondShake

	time.Sleep(time.Second)

	thirdShake := <-comms
	ackS = thirdShake.seq + 1
	fmt.Printf("ackS is %d \n", ackS)
	fmt.Printf("seqS is %d \n\n", thirdShake.ack)
}

func main() {

	var comms = make(chan Data, 2)

	go client(comms)
	go server(comms)

	time.Sleep(2 * time.Second)
}
