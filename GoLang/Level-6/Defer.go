package main

import "fmt"

func main () {
  // EDIT HERE
  defer foo()
  bar()
}

func foo() {
  fmt.Println("foo")
}

func bar() {
  fmt.Println("bar")
}
