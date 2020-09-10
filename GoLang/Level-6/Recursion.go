package main

import "fmt"

func main () {
  // EDIT HERE
  fmt.Println(4 * 3 * 2 * 1)

  n1 := factorial(4)
  n2 := loop(4)
  fmt.Println(n1)
  fmt.Println(n2)
}

func factorial(n int) int {
  if n == 0 {
    return 1
  }
    return n * factorial(n-1)
}

func loop(n int) int {
  t := 1
  for; n > 0; n-- {
    t *= n
  }
  return t
}
