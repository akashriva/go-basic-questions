package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countWords("I love programming in Go"))
}

func countWords(str string) int{
	words := strings.Fields(str)
	count := len(words)
	return count
	
}
