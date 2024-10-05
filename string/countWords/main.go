/*
1. Count Words in a String
Write a function CountWords that takes a string as input and returns the number of words in the string. A word is defined as a sequence of characters separated by spaces.
("The quick brown fox")) // 4
*/
/*  strings.Fields :- fields function is used split the with white space and consecutive space. and return the slice of sub string and
or empty slice if string is only with space
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countWords("The quick brown fox"))
}

func countWords(str string) int {
	words := strings.Fields(str)
	count := len(words)
	return count

}
