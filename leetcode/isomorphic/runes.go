package main

import "fmt"

func main() {
	s := "badc"
	t := "baba"
	fmt.Println(isIsomorphic(s, t))
}

func isIsomorphic(s string, t string) bool {
	slen, tlen := len(s), len(t)
	if slen != tlen {
		return false
	}

	sRunes := []rune(s)
	tRunes := []rune(t)
	m := make(map[rune]rune)
	d := make(map[rune]bool)

	for i := 0; i < len(sRunes); i++ {
		if exist, ok := m[sRunes[i]]; ok {
			if exist != tRunes[i] {
				return false
			}
			continue
		}

		if _, ok := d[tRunes[i]]; ok {
			return false // catch duplicate runes in m
		}

		m[sRunes[i]] = tRunes[i]
		d[tRunes[i]] = true
	}

	return true
}
