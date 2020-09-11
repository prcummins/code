package main

import "fmt"

func main () {
// EDIT HERE
  a := foo()
  fmt.Println(a())
}

func foo () func () int {
  return func() int {
    return 42
  }
}
