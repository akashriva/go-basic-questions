/*
3. Title Case Conversion
Write a function ToTitleCase that takes a string as input and returns the string with the first letter of each word capitalized.
fmt.Println(ToTitleCase("hello world"))   // "Hello World"


unicode.ToUpper() is used to convert the character to uppercase and its take character as argument and return character in uppercase
string.ToLower take string as argument and return string in lower case 
*/
package main

import(
	"fmt"
	"strings"
	"unicode"
)

func main(){
	fmt.Println(titleCase("hello sachin bhai"))
}

func titleCase(str string) string{
	words := strings.Fields(str)

	for i, word := range words{
		if len(word) > 0 {
			first := string(unicode.ToUpper(rune(word[0]))) 
			fmt.Println(first)
			rest := strings.ToLower(word[1:])               
			words[i] = first + rest
		}
	}

	return strings.Join(words, " ")
}