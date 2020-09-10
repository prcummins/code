package main

import "fmt"

func main () {
  // EDIT HERE
  i := []int{1,2,3,4,5,6,7,8,9}

  s := sum(i...)
  fmt.Println("All numbers summed are equal to:", s)

  e := even(sum, i...)
  fmt.Println("All even numbers summed are equal to:", e)

  o := odd(sum, i...)
  fmt.Println("All odd numbers summed are equal to:", o)
}

func sum (xi...int) int {
  t := 0
  for _, v := range xi {
    t += v
  }
  return t
}

func even (f func(xi...int)int, xxi...int) int {
  var yi []int
  for _, v := range xxi {
    if v%2 == 0 {
      yi = append(yi, v)
    }
  }
  return f(yi...)
}

func odd (f func(xi...int)int, xxi...int) int {
  var yi []int
  for _, v := range xxi {
    if v%2 != 0 {
      yi = append(yi, v)
    }
  }
  return f(yi...)
}
