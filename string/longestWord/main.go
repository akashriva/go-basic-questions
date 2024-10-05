/*
2. Find the Longest Word
Write a function LongestWord that takes a string as input and returns the longest word in the string. If there are multiple words with the same maximum length, return the first one.
("I love programming in Go")) // "programming"
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestWord("I love programming in Go"))
}

func longestWord(str string) string {
	strSlice := strings.Fields(str)
	longest := ""
	for _, word := range strSlice {
		if len(word) > len(longest) {
			longest = word
		}
	}

	return longest
}
