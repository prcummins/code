package main

import "fmt"

func main () {
  // EDIT HERE
  s1 := foo()
  fmt.Println(s1)

  s2 := bar()
   i := s2()
   fmt.Println(i)
}

func foo () string {
  s := "Hello World!"
  return s
}

func bar() func() int {
  return func() int {
    return 451
  }
}
