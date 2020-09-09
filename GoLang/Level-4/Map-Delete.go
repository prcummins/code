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
  fmt.Println(m["JC"])

  delete(m, "PC")
  fmt.Println(m)

  if v, ok := m["JC"]; ok {
    fmt.Println("Value:", v)
  }

}
