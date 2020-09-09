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
  fmt.Println(p1.First, p1.Last)
  for i1, v1 := range p1.IC {
    fmt.Println(i1, v1)
  }
  fmt.Println(p2.First, p2.Last)
  for i2, v2 := range p2.IC {
    fmt.Println(i2, v2)
  }
}
