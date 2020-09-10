package main

import "fmt"

type person struct {
  First   string
  Last    string
  IC      []string
}

func main () {
  // EDIT HERE
  p1 := person {
    First:    "James",
    Last:     "Bond",
    IC:       []string{"Vanilla", "Chocolate", "Maple"},
  }
  p2 := person {
    First:    "Paul",
    Last:     "Cummins",
    IC:       []string{"Chocolate", "Oreo", "Snickers"},
  }

  m := map[string]person{
    p1.Last:  p1,
    p2.Last:  p2,
  }
  for i3, v3 := range m {
    fmt.Println(i3)
      for i4, v4 := range v3.IC{
        fmt.Println(i4, v4)
      }
  }
}
