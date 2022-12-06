// count items in a slice (recursion)
package main

import "fmt"

func main() {
	list := []string{"first", "second", "third", "forth"}
	fmt.Println(count(list))
}

func count(list []string) int {
	if len(list) == 0 {
		return 0
	}
	return 1 + count(list[1:])
}
