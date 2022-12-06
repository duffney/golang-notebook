package main

import (
	"fmt"
)

func main() {
	s := "axc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 { // base case
		return true
	}
	var index, count int

	// loop through t, if there a match increase index to move to the next letter of s, use a counter to
	// return true when the length of s is reached
	for i := 0; i < len(t); i++ {
		if t[i] == s[index] { // if the byte representation of the char matches
			count++ // increase counter
			index++ // increase index
		}
		if count == len(s) { // prevent out of range index error return if len of s reached
			return true
		}
	}
	return false
}
