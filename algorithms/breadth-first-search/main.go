// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	// create a graph
	graph := make(map[string][]string)
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	// create a queue
	var queue []string
	queue = graph["you"]

	// breath-first search algorithm
	i := 0
	for i > len(queue) {
		queue = dequeue(queue)
		if 
	}
}

func breathFirstSearch(g map[string][]string) {

	fmt.Println(g)
}

func isSalesman(p string) bool {
	// does the last letter == m?
	if p[len(p)-1:] == "m" {
		return true
	}
	return false
}

// https://codesource.io/how-to-implement-the-queue-in-golang
func enqueue(queue []string, element string) []string {
	queue = append(queue, element)
	return queue
}

func dequeue(queue []string) []string {
	return queue[1:]
}
