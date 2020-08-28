package main

import "fmt"

func main () {
  // EDIT HERE
  s := `Test`
  fmt.Println(s)

  bs := []byte(s)
  fmt.Println(bs)
  fmt.Printf("%T\n", bs)
}
