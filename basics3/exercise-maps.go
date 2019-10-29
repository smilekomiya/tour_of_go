package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	var wordsCount = map[string]int{}
	for _, v := range words {
		count := wordsCount[v]
		wordsCount[v] = count + 1
	}

	return wordsCount
}

func main() {
	// fmt.Println(WordCount("hoge fuga"))
	wc.Test(WordCount)
}
