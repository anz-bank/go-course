# Lab 3 - Letter frequency

- Create an executable go program in directory `03_letters/USERNAME`
- Write a function that returns a mapping of each letter to its frequency:

```
func letters(s string) map[rune]int
```

Write a function that returns a sorted slice of strings with elements `"{key}:{val}"`. Use package [sort](https://golang.org/pkg/sort/):

```
func sortLetters(m map[rune]int) []string
```

Call `fmt.Println(strings.Join(sortLetters(letters("aba")), "\n"))` in `main` to print:

```
a:2
b:1
```

Bonus points: comprehensive tests
