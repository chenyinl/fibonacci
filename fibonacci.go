package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
   last  := 0
   last1 := 0
   last2 := 1
   return func() int {
       last = last1 + last2
	   last1 = last2
	   last2 = last
	   return last
   }
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
