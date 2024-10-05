package main

import (
	"fmt"
)

func longestConsecutiveSubsequence(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	// Create a set of all elements in the array for fast lookup
	elementSet := make(map[int]bool)
	for _, num := range arr {
		elementSet[num] = true
	}

	longestStreak := 0

	// Iterate through the array and find the start of a sequence
	for _, num := range arr {
		// Check if the current number is the start of a sequence
		if !elementSet[num-1] { // num-1 should not exist in the set, indicating the start of a sequence
			currentNum := num
			currentStreak := 1

			// Continue the sequence by checking for the next consecutive numbers
			for elementSet[currentNum+1] {
				currentNum++
				currentStreak++
			}

			// Update the maximum length found
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}

	return longestStreak
}

func main() {
	array := []int{100, 4, 200, 1, 3, 2}
	result := longestConsecutiveSubsequence(array)
	fmt.Printf("The length of the longest consecutive subsequence is: %d\n", result) // Output: 4 (sequence is 1, 2, 3, 4)
}
