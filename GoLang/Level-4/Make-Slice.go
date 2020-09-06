package main

import "fmt"

func main () {
  // EDIT HERE
  x := make([]int, 10, 12)
  fmt.Println(x)
  fmt.Println(len(x))
  fmt.Println(cap(x))

  x = append(x, 11)
  x = append(x, 12)
  x = append(x, 13)
  fmt.Println(x)
  fmt.Println(len(x))
  fmt.Println(cap(x))

}
