package main

import "fmt"

func main () {
  // EDIT HERE
  x := []int{1, 2, 3}
  y := []int{4, 5, 6}
  x = append(x, y...)
  fmt.Println(x)
  x = append(x[:1],x[2:]...)
  fmt.Println(x)
}
