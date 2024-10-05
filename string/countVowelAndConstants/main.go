/*
4. Count Vowels and Consonants
Write a function CountVowelsAndConsonants that takes a string as input and returns two integers: the number of vowels and the number of consonants in the string.
 Consider only alphabetic characters and ignore case.
vowels, consonants := CountVowelsAndConsonants("Hello, World!")
fmt.Println(vowels, consonants) // 3 7
*/
/* Unicode.IsLetter() is take character as argument if he value is Alphabetical character is return true else false
strings.ContainsRune its take strings as first argument and second as character if the character find in the string return true
else false
*/
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(countVowelsConstants("stings"))
}

func countVowelsConstants(str string) (vowel int, constants int) {
	vowels := "AaEeIiOoUu"

	for _, char := range str {
		if unicode.IsLetter(char) {
			if strings.ContainsRune(vowels, char) {
				vowel++
			} else {
				constants++
			}
		}
	}

	return vowel, constants
}
