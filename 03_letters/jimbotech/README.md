# Lab3 Letter Sorter
Implementation of lab 3 from ANZs go course
## Overview
Two functions support the required functionality:

    func letters(s string) map[rune]int

Which counts the occurrence of the individual letter in the input string s and produces a map of those runes (characters) and the count of their occurrence. This function practises the use of maps.

    func sortLetters(m map[rune]int) []string

Takes the map with key=runes and value=integer and creates a slice of strings in the form of "character:integer" in a sorted order. This function demonstrates the creating of strings from runes and numbers, appending to a slice and using the go function sort.

## Funky Features
Strictly speaking I should have creates the slice of the keys, sorted it and then iterate over the sorted slice, retrieve the associated map value and constructed the slice of strings. However, I took a shortcut by creating the string first and then sorting that slice. It is more efficient and less work and I just hope there are no cases where such a cunning plan might fail.

> Written with [StackEdit](https://stackedit.io/).
<!--stackedit_data:
eyJoaXN0b3J5IjpbMTEzNjE2NDU4MCw3MzA5OTgxMTZdfQ==
-->