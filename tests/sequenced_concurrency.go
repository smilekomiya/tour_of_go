package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str   string
	block chan int
}

func main() {
	c := fanIn(generator("Hello"), generator("Bye"))
	for i := 0; i < 10; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)

		msg2 := <-c
		fmt.Println(msg2.str)

		<-msg1.block // reset channel, stop blocking
		<-msg2.block
	}
}

func fanIn(c1, c2 <-chan Message) <-chan Message {
	new_c := make(chan Message)
	go func() {
		for {
			new_c <- <-c1
		}
	}()
	go func() {
		for {
			new_c <- <-c2
		}
	}()
	return new_c
}

func generator(msg string) <-chan Message {
	c := make(chan Message)
	blockingStep := make(chan int)
	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", msg, i), blockingStep}
			// time.Sleep(time.Second)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			blockingStep <- 1 // blocking by waiting for input
		}
	}()

	return c
}
