package main

import "fmt"

type person struct {
  first   string
  last    string
  age     int
}

func main () {
  // EDIT HERE
  p1 := person {
    first:  "James",
    last:   "Bond",
    age:    32,
  }
  p1.speak()
}

func (p person) speak () {
  fmt.Println("Hello my name is", p.last,",", p.first, p.last,",", "and I am", p.age)
}
