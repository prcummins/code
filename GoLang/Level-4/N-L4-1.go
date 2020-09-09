package main

import "fmt"

func main () {
  // EDIT HERE
  a := [5]int{5, 6, 7, 8, 9}
  fmt.Println(a)
    for i, v := range a {
      fmt.Println(i, v)
    }
    fmt.Printf("%T\n", a)
}
