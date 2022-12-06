package main

import "fmt"

func main() {
	set := make(map[string]struct{})
	var exists = struct{}{}
	statesNeeded := []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}
	for _, v := range statesNeeded {
		set[v] = exists
	}

	for state := range set {
		fmt.Println(state)
	}

	//stations = {}
	//stations["kone"] = set(["id", "nv", "ut"])
	//stations["ktwo"] = set(["wa", "id", "mt"])
	//stations["kthree"] = set(["or", "nv", "ca"])
	//stations["kfour"] = set(["nv", "ut"])
	//stations["kfive"] = set(["ca", "az"])
}
