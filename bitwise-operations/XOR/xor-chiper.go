// https://www.educative.io/blog/bit-manipulation-algorithms#xor
// TODO: Write a simple XOR cipher block algorithm
package main

import (
	"fmt"
)

func main () {
	// cipher 1 as 3 with XOR
	plaintext := 1 // 00000001
	key := 2 // 00000010

	fmt.Println("simple...")
	ciphertext := plaintext ^ key// 00000011 
	//compares two bits and returns 1 if those bits are different. 
	//If both bits are the same, then XOR returns a 0.
	//allows you to compre two bits and generate a completely different bit
	fmt.Printf("encrypt: %d ^ %d = %d\n", plaintext, key, ciphertext)

	xorOne := ciphertext^ key 
	fmt.Printf("decrypt: %d ^ %d = %d\n", ciphertext, key, xorOne)
	fmt.Println(fmt.Sprintf("%08b",xorOne))

	a := "a"
	binA := fmt.Sprintf("%08b", a[0])
	fmt.Println(binA)
	fmt.Printf("Character: %s, Byte(ascii): %d, Binary: %s\n", a, a[0], binA)
}

func encrypt (plaintext, key string) string {
	cipherText := make([]byte, len(plaintext))
	keyLen := len(key)

	for i := range plaintext {
		cipherText[i] = plaintext[i] ^ key[i % keyLen]
		fmt.Println(key[i % keyLen])
	}

	return string(cipherText)
}

func decrypt (s string) string {
	return ""
}
