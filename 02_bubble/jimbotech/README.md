# Bubble Sort - lab 2 of the go course
## Details
Implements a function that does a bubble sort, which keeps iterating over the collection and swaps out of order elements. I presume the key to this exercise it to discover the funky tuple assignment in go.
  

     a, b = b, a

Now this probably does not make any sense to normal humans, but it does in go, as it has the construct a, b = func(), which means func returns two values that are then assigned to the variables as expected. When considering this, the above statement all of a sudden does compute.
The second aim, I guess, is to use the "range" a lot instead of traditional for loops. All this can be found in my implementation in 

    func bubble(s []int) []int


> Written with [StackEdit](https://stackedit.io/).
