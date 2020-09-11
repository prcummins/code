package main

import "fmt"

func main () {
  // EDIT HERE
  v1 := []int{1,2,3,4,5}
  n1 := foo (v1...)
  defer fmt.Println(n1)

  v2 := []int{1,2,3,4,5,6}
  n2 := bar (v2)
  fmt.Println(n2)
}

func foo (x ...int) int {
  t := 0
  for _, v := range x {
    t += v
  }
  return t
}

func bar (x []int) int {
  t := 0
  for _, v := range x {
    t += v
  }
  return t
}
