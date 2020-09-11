package main

import "fmt"

func main () {
  // EDIT HERE
  a := foo ()
  b, c := bar ()

  fmt.Println(c, a, "and on September 20th, I will be", b)
}

func foo() int {
  return 42
}

func bar() (int, string) {
  return 43, "I am"
}
