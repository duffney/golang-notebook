// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	s := "foo"
	t := "bar"
	fmt.Println(isIsomorphic(s, t))
}

func isIsomorphic(s string, t string) bool {
	slen, tlen := len(s), len(t)
	if slen != tlen {
		return false
	}
	// convert to runes instead|
	mapping := make(map[byte]byte)   // store mapping s => t
	apperance := make(map[byte]bool) // store if character already appears on t
	for i := 0; i < slen; i++ {
		fmt.Println(s[i])
		fmt.Println(mapping[s[i]])
		if existing, ok := mapping[s[i]]; ok { // does the byte value of the character already exist in the map?
			fmt.Println(existing, t[i])
			if existing != t[i] { // does existing not match the byte value of t at the same position?
				return false // the existing key has two possible values and therefore cannot be isomoprhic
			}
			continue
		}

		if _, ok := apperance[t[i]]; ok {
			return false // why?
		}
		mapping[s[i]] = t[i]
		apperance[t[i]] = true
	}
	return true
}
