package main

import "fmt"

type car struct {
  Color  string
  Doors  string
}

type model struct {
  car
  Model  string
}

func main () {
  // EDIT HERE
  c1 := model {
      car: car {"Red", "4",},
      Model: "Porche",
  }
  fmt.Println(c1.Model)
  c1.vehicle()
}

func (c model) vehicle () {
  fmt.Println("You chose a", c.Color, c.Doors, "Door", c.Model)
}
