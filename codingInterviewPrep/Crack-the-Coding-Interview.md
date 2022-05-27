# Crack the Coding Interview

## Arrays and Strings

> **Is Unique**: Implement an algorithm to determine if a string has all unique characters.

Soulution: https://play.golang.com/p/zfqzmMoZ1VR

---

> **Check permutations**: Given two strings, write a method to detect if one is a permutation of the other.

solution: https://play.golang.com/p/cwbgH2FmJzg

---

> **URLify**: Write a method to replace all spaces in a string with `%20`. You may assume that the string has sufficient space at the end to hold the additional characters, and that you are given the "true" length of the string.

Input: "Mr John Smith    " (13 chars)

Output: `Mr%20John%20Smith`

Brue-force:
	- Use strings.trim to remove whitespace at the end and strings.replace to replace whitespace with `%20`.

solution: https://play.golang.com/p/DBUrbwT6Hcq

---

> **Palindrome Permutation**: Given a string, write a function to check if it is a permutation of a palindrome.

Example
	- Input: Tact Coa
	- Output: True (permutations: "taco cat", atco cta", etc.)

Partial Solution: https://play.golang.com/p/zaZAx_wz6-J

---

> **One Away**: There are three types of edits that can be performed on strings: insert a character, remove a character, or replace a character. Given two strings, write a function to check if they are one edit (or zero edits) away.

Example:
- pale, ple -> true
- pales, pale -> true
- pale, bale -> true
- pale, bake -> false

Partial solution: https://play.golang.com/p/Tus2QIvwtK-