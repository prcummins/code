package main

import "fmt"

func main () {
  // EDIT HERE
  s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
  s = append(s[:3], s[6:]...)
  fmt.Println(s)
}
