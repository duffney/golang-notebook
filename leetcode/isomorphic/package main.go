package main

import "fmt"

func main() {
	s := "egg"
	t := "add"
	fmt.Println(isIsomorphic(s, t))
}
func isIsomorphic(s string, t string) bool {
	slen, tlen := len(s), len(t)
	if slen != tlen {
		return false
	}

	mapping := make(map[byte]byte)   // store mapping s => t
	apperance := make(map[byte]bool) // store if character already appears on t
	for i := 0; i < slen; i++ {
		if existing, ok := mapping[s[i]]; ok {
			if existing != t[i] {
				return false
			}
			continue
		}

		if _, ok := apperance[t[i]]; ok {
			return false
		}
		mapping[s[i]] = t[i]
		apperance[t[i]] = true
	}
	return true
}
