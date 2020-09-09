package main

import "fmt"

func main () {
  // EDIT HERE
  s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
  sa1 := []int{52}
  s = append(s, sa1...)
  fmt.Println(s)
  sa2 := []int{53, 54, 55}
  s = append(s, sa2...)
  fmt.Println(s)
}
