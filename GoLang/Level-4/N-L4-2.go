package main

import "fmt"

func main () {
  // EDIT HERE
  s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  fmt.Println(s)
    for i, v := range s {
      fmt.Println(i, v)
    }
    fmt.Printf("%T\n", s)
}
