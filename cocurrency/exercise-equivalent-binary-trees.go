package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// type Tree struct {
//     Left  *Tree
//     Value int
//     Right *Tree
// }

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		walk(t.Right, ch)
	}
}

func Same(tree1, tree2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(tree1, ch1)
	go Walk(tree2, ch2)

	var tree1_slice []int
	var tree2_slice []int

	for v := range ch1 {
		tree1_slice = append(tree1_slice, v)
	}

	for v := range ch2 {
		tree2_slice = append(tree2_slice, v)
	}

	for i := 0; i < 10; i++ {
		if tree1_slice[i] != tree2_slice[i] {
			return false
		}
	}

	return true
}

func main() {
	// tree := tree.New(1)
	// ch := make(chan int)
	// go Walk(tree, ch)
	//
	// for v := range ch {
	// 	fmt.Println(v)
	// }

	tree1 := tree.New(1)
	tree2 := tree.New(2)

	fmt.Println(Same(tree1, tree1))
	fmt.Println(Same(tree1, tree2))
	fmt.Println(Same(tree2, tree2))
}
