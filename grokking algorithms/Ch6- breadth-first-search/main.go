// Credit: Grokking Algorithms: An illustrated guide for programmers and other curious people
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

	fmt.Println(breathFirstSearch(graph))
}

func breathFirstSearch(graph map[string][]string) string {
	// create a queue
	var queue []string
	var person string
	queue = graph["you"]

	i := 0
	for i < len(queue) { // while loop
		person, queue = dequeue(queue)
		if isSalesman(person) {
			return person
		} else {
			for _, contact := range graph[person] {
				queue = enqueue(queue, contact)
			}
		}
	}
	return ""
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

func dequeue(queue []string) (string, []string) {
	return queue[0], queue[1:]
}
