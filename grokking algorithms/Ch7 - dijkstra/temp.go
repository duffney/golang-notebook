package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	names := []string{"Ana", "Bob", "Clair", "Daniel", "Fred", "Hush", "Timmy"}

	// some string to check
	nameCheck := "Bob"
	isPresent1 := slices.Contains(names, nameCheck)

	if isPresent1 {
		fmt.Println(nameCheck, "is in the names array")
	} else {
		fmt.Println(nameCheck, "is not in the names array")
	}

	scores := []float64{14.2, 26.5, 9.6, 36.4, 52.6}
	scoreCheck := 9.6

	isPresent2 := slices.Contains(scores, scoreCheck)
	if isPresent2 {
		fmt.Println(scoreCheck, "is in the scores array")
	} else {
		fmt.Println(scoreCheck, "is not in the scores array")
	}
}
