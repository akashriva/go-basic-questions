/*
Check if a String Starts with a Prefix
Write a function HasPrefix that takes two strings as input: str and prefix. The function should return a boolean indicating whether str starts with prefix.

strings.HashPrefix(sring, prefix) is take two argument first as a string and second as string the prefix and its return boolen value
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	prefix := "Go"
	str := "Golang"
	result := checkPrefix(str, prefix)
	fmt.Println(result)

}

func checkPrefix(str string, prefix string) bool {

	if strings.HasPrefix(str, prefix) {
		return true
	} else {
		return false

	}

}
