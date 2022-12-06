package main

import (
	"fmt"
	"math"

	"golang.org/x/exp/slices"
)

func main() {
	//https://nilpointer.net/programming-language/golang/nested-maps-in-golang/
	graph := make(map[string]map[string]interface{})

	graph["start"] = make(map[string]interface{})
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2

	graph["a"] = make(map[string]interface{})
	graph["a"]["fin"] = 1

	graph["b"] = make(map[string]interface{})
	graph["b"]["a"] = 3
	graph["b"]["fin"] = 5

	graph["fin"] = make(map[string]interface{})

	//costs
	infinity := math.Inf(1)
	costs := make(map[string]int)
	costs["a"] = 6
	costs["b"] = 2
	costs["fin"] = int(infinity)

	// parents
	parents := make(map[string]string)
	parents["a"] = "start"
	parents["b"] = "start"
	parents["fin"] = ""

	// array of processed nodes
	var processed []string

	node := findLowestCostNode(costs, processed)
	for node != "" {
		cost := costs[node]
		neighbors := graph[node]
		for n := range neighbors {
			newCost := cost + neighbors[n].(int)
			if costs[n] > newCost {
				costs[n] = newCost
				parents[n] = node
			}
		}
		processed = append(processed, node)
		node = findLowestCostNode(costs, processed)
	}

	fmt.Println(parents)
}

func findLowestCostNode(costs map[string]int, process []string) string {
	lowestCost := 1000
	var lowestCostNode string

	for node := range costs {
		cost := costs[node]
		processed := slices.Contains(process, node)

		if (cost < lowestCost) && (!processed) {
			lowestCostNode = node
			lowestCost = cost
		}
	}
	return lowestCostNode
}
