package main

import "fmt"
import "math"

type SQUARE struct {
  length  float64
  width   float64
}

type CIRCLE struct {
  radius float64
}

type shape interface {
  area () float64
}

func main () {
  circ := CIRCLE {
    radius: 9.9,
  }
  squa := SQUARE {
    length: 9,
    width:  9,
  }
  info (circ)
  info (squa)
}

func (s SQUARE) area () float64 {
  return s.length * s.width
}

func (c CIRCLE) area () float64 {
  return math.Pi * c.radius * c.radius
}

func info (s shape) {
  fmt.Println(s.area())
}
