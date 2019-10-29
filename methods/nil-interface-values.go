package main

import "fmt"

type I interface {
	M()
}

type Hoge string

func (h Hoge) M() {
	fmt.Println(h)
}

func main() {
	var i I
	i = Hoge("hoge")

	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
