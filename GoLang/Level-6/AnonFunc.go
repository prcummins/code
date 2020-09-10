package main

import "fmt"

func main () {
  // EDIT HERE
 foo ()

 func() {
      fmt.Println("This is an annonymous function")
   }()

 func(x int) {
      fmt.Println("The best age is age:", x)
   }(42)

}

func foo () {
  fmt.Println("This function was called from main")
}
