package main

import "fmt"

func fibonacci() func() int {
	var now int
	a := 0
	b := 1
	return func() int {
		now = a
		a, b = b, a+b
		return now
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
