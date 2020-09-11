package main

import "fmt"

func main () {
  // EDIT HERE
  a := incrementor()
  b := incrementor()
  fmt.Println(a())
  fmt.Println(a())
  fmt.Println(b())
  fmt.Println(a())
}

func incrementor () func () int {
  var x int
  return func () int {
    x++
    return x
  }
}
