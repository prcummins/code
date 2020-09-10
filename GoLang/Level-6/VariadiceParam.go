package main

import "fmt"

func main () {
  // EDIT HERE
    foo (1, 2, 3, 4)
}

func foo (x...int) {
    fmt.Println(x)
    fmt.Printf("%T\n", x)
    t := 0
    for _, v := range x {
      t += v
    }
    fmt.Println(t)
}
