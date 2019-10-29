package main

import (
	"fmt"
	"math/rand"
)

func boring(msg string, quit <-chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case c <- fmt.Sprintf("%s", msg):
				// do nothing
			case <-quit:
				return
			}
			// c <- fmt.Sprintln(msg)
			// time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	quit := make(chan bool)
	c := boring("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
}
