# Lab 5 - Stringer

- Create an executable go program in directory `05_stringer/USERNAME`
- Make the `IPAddr` type implement `fmt.Stringer` to print the address as a dotted quad
- Find hints at [tour of go exercise: stringers](https://tour.golang.org/methods/18)
- Call `fmt.Println(IPAddr{127, 0, 0, 1})` in `main` to print:

```
127.0.0.1
```
