package main

import "fmt"

func main () {
  // EDIT HERE
  m := map[string]int{
      "PC": 42,
      "JC": 13,
  }
  fmt.Println(m)
  fmt.Println(m["PC"])
  fmt.Println(m["WC"])

  v, ok := m["WC"]
  fmt.Println(v)
  fmt.Println(ok)

  if v, ok := m["JC"]; ok {
    fmt.Println("THIS WORKED!", v)
  }
}
