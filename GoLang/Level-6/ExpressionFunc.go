package main

import "fmt"

func main () {
  // EDIT HERE
  fmt.Println("Learning how to use func expressions")

  f1 := func () {
      fmt.Println("My first func expression")
  }
  f1()

  f2 := func (x int) {
      fmt.Println("A func expression with a required argument =", x)
  }
  f2(1984)
}
